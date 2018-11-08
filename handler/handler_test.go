package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/config"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/httptest"
)

func init() {
	err := config.Init("../conf/config_test.yaml")
	if err != nil {
		panic(err)
	}

	err = model.Init()
	if err != nil {
		panic(err)
	}

	g := gin.New()
	Register(g)

	httptest.GinEngine = g
}

func SignUpAndLogin(username string) (token string, err error) {

	//sign up first
	_, err = DoSignUp(goutil.Map{"username": username, "password": "123"})
	if err != nil {
		return
	}
	res, err := DoLogin(username, "123")
	if err != nil {
		return
	}
	token = res.GetString("token")
	return
}

func removeAll() {
	model.DB.C(model.CollUser).RemoveAll(nil)
	model.DB.C(model.CollSeq).RemoveAll(nil)
	model.DB.C(model.CollOp).RemoveAll(nil)
	model.DB.C(model.CollTask).RemoveAll(nil)
	model.DB.C(model.CollProject).RemoveAll(nil)
	model.DB.C(model.CollTag).RemoveAll(nil)
}
