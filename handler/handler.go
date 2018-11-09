package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/hzxiao/taskmeter/router/middleware"
	"net/http"
	"strconv"
	"strings"
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

type Arg struct {
	Key          string
	Value        interface{}
	DefaultValue string
	Type         string
	Require      bool
}

func CheckURLArg(formValue map[string][]string, args []*Arg) (goutil.Map, error) {
	argMap := goutil.Map{}
	if len(args) == 0 {
		return argMap, nil
	}

	for _, arg := range args {
		vs := formValue[arg.Key]
		var v string
		if len(vs) == 0 {
			if arg.Require {
				return nil, errno.New(errno.ErrQueryValueInvalid, fmt.Errorf("require %v field", arg.Key))
			}
			if arg.DefaultValue != "" {
				v = arg.DefaultValue
			} else {
				continue
			}
		} else {
			v = vs[0]
		}
		if v == "" {
			continue
		}
		switch arg.Type {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint32", "uint64":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, errno.New(errno.ErrQueryValueInvalid,
					fmt.Errorf("convert to int64 err(%v) by key(%v), value(%v)", err, arg.Key, v))
			}
			argMap.Set(arg.Key, i)
		case "float32", "float64":
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, errno.New(errno.ErrQueryValueInvalid,
					fmt.Errorf("convert to float64 err(%v) by key(%v), value(%v)", err, arg.Key, v))
			}
			argMap.Set(arg.Key, f)
		case "bool":
			v = strings.ToLower(v)
			if v == "1" || v == "true" {
				argMap.Set(arg.Key, true)
			} else if v == "0" || v == "false" {
				argMap.Set(arg.Key, false)
			} else {
				return nil, errno.New(errno.ErrQueryValueInvalid,
					fmt.Errorf("invalid bool type of key(%v), value(%v)", arg.Key, v))
			}
		case "string":
			if v != "" {
				argMap.Set(arg.Key, v)
			}
		default:
			return nil, errno.New(errno.ErrQueryValueInvalid,
				fmt.Errorf("unknown type(%v) of key(%v)", arg.Type, arg.Key))
		}
	}

	return argMap, nil
}

func formatSort(sortBy, sortModel string) []string {
	if sortModel == "desc" {
		sortBy = "-" + sortBy
	}

	return []string{sortBy}
}