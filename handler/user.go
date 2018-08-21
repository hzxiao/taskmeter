package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/errno"
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

	user, err := model.InsertUser(data)
	if err == nil {
		log.Infof("[SignUp] sign up by user(%v)", goutil.Struct2Json(user))
		SendResponse(c, nil, goutil.Map{"user": user})
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

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Hello(c *gin.Context)  {
	SendResponse(c, nil, "hello")
}

