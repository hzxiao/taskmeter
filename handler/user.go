package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/httptest"
	"github.com/hzxiao/taskmeter/pkg/token"
	"github.com/hzxiao/taskmeter/service"
	"github.com/lexkong/log"
	"strings"
)

// @Summary SignUp a new user
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /signup [post]...
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
	return checkResultError(httptest.PostJSON("/signup", data))
}

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
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
		ID:       user.GetString("id"),
		Username: user.GetString("username"),
	}, "")
	if err != nil {
		log.Errorf(err, "[Login] generate token by data(%v)", goutil.Struct2Json(user))
		SendResponse(c, errno.New(errno.InternalServerError, err).Add("generate token"), nil)
		return
	}

	err = service.AddSignInRecord(user.GetString("id"), goutil.Map{
		"username": user.GetString("username"),
		"ip":       c.ClientIP(),
	})
	if err != nil {
		log.Errorf(err, "[Login] add login record by uid(%v)", user.GetString("id"))
	}

	SendResponse(c, nil, goutil.Map{"token": tokenString})
}

func DoLogin(username, password string) (goutil.Map, error) {
	return checkResultError(httptest.PostJSON("/login", goutil.Map{"username": username, "password": password}))
}
