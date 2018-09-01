package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"net/http"
	"github.com/hzxiao/goutil"
	"fmt"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Register(g *gin.Engine)  {
	g.POST("/signup", SignUp)
	g.POST("/login", Login)
}

func checkResultError(data goutil.Map, err error) (goutil.Map, error) {
	if err != nil {
		return nil, err
	}

	if data.GetInt64("code") > 0 {
		return nil, fmt.Errorf("err- code: %v, message: %v", data.GetInt64("code"), data.GetString("message"))
	}

	return data.GetMap("data"), nil
}