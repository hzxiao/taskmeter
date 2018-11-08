package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/hzxiao/taskmeter/router/middleware"
	"net/http"
)

var (
	StartMoment = timeutil.Now()
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

func Register(g *gin.Engine) {
	v1 := g.Group("/api/v1")

	pub := v1.Group("/pub")
	pub.GET("/ping", Ping)
	pub.POST("/signup", SignUp)
	pub.POST("/login", Login)

	usr := v1.Group("/usr")
	usr.Use(middleware.Auth())

	//tasks
	usr.POST("/projects/:project_id/tasks", AddTask)
	usr.GET("/projects/:project_id/tasks", ListTasks)
	usr.PUT("/tasks/:task_id", UpdateTask)
	usr.GET("/tasks/:task_id", LoadTask)
	usr.DELETE("/tasks/:task_id", DelTask)
	usr.GET("/search/tasks", SearchTasks)
}

func Ping(c *gin.Context) {
	SendResponse(c, nil, goutil.Map{
		"start":       StartMoment,
		"startFormat": timeutil.GetDateString(StartMoment),
		"now":         timeutil.Now(),
		"nowFormat":   timeutil.GetDateString(timeutil.Now()),
	})
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

func newArgInvalidError(format string, message ...interface{}) error {
	err := errno.New(errno.ErrApiArgumentInvalid, nil)
	err.Addf(format, message...)
	return err
}
