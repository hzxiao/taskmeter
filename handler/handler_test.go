package handler

import (
	"github.com/gin-gonic/gin"
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
