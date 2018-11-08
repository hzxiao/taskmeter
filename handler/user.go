package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/httptest"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/hzxiao/taskmeter/pkg/token"
	"github.com/hzxiao/taskmeter/service"
	"github.com/lexkong/log"
	"strconv"
	"strings"
)

func SignUp(c *gin.Context) {
	var data goutil.Map
	if err := c.Bind(&data); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	data.Set("method", constvar.UsernameRegister)
	user, err := service.SignUp(data)
	if err == nil {
		log.Infof("[SignUp] sign up by user(%v)", goutil.Struct2Json(user))
		SendResponse(c, nil, goutil.Map{"username": data.GetString("username")})
		return
	}

	if strings.Contains(err.Error(), "uname dup key") {
		log.Debugf("[SignUp] sign up by data(%v) which username already exists", goutil.Struct2Json(data))
		SendResponse(c, errno.ErrUsernameExist, nil)
		return
	}

	log.Errorf(err, "[SignUp] sign up by data(%v)", goutil.Struct2Json(data))
	SendResponse(c, err, nil)
}

func DoSignUp(data goutil.Map) (goutil.Map, error) {
	return checkResultError(httptest.PostJSON("/api/v1/pub/signup", data))
}

func Login(c *gin.Context) {
	var data goutil.Map
	if err := c.Bind(&data); err != nil {
		log.Errorf(err, "[Login} bind data")
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := service.Login(data.GetString("username"), data.GetString("password"))
	if err != nil {
		log.Errorf(err, "[Login] login by data(%v)", goutil.Struct2Json(data))
		SendResponse(c, err, nil)
		return
	}

	tokenString, err := token.GenerateToken(token.Context{
		ID:           user.GetString("id"),
		Username:     user.GetString("username"),
		GenerateTime: strconv.FormatInt(timeutil.Now(), 10),
		Source:       constvar.LoginFromWeb,
	}, "")
	if err != nil {
		log.Errorf(err, "[Login] generate token by data(%v)", goutil.Struct2Json(user))
		SendResponse(c, errno.New(errno.InternalServerError, err).Add("generate token"), nil)
		return
	}

	err = service.AddSignInRecord(user.GetString("id"), goutil.Map{
		"username": user.GetString("username"),
		"ip":       c.ClientIP(),
		"source":   constvar.LoginFromWeb,
	})
	if err != nil {
		log.Errorf(err, "[Login] add login record by uid(%v)", user.GetString("id"))
	}

	SendResponse(c, nil, goutil.Map{"token": tokenString})
}

func DoLogin(username, password string) (goutil.Map, error) {
	return checkResultError(httptest.PostJSON("/api/v1/pub/login", goutil.Map{"username": username, "password": password}))
}
