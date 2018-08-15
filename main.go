package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/taskmeter/config"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/version"
	"github.com/hzxiao/taskmeter/router"
	"github.com/hzxiao/taskmeter/router/middleware"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"net/http"
)

var (
	cfg = pflag.StringP("config", "c", "", "taskmeter config file path.")
	v   = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *v {
		err := version.Print()
		if err != nil {
			panic(err)
		}
		return
	}

	//init config
	err := config.Init(*cfg)
	if err != nil {
		panic(err)
	}

	//init db
	err = model.Init()
	if err != nil {
		panic(err)
	}

	g := gin.New()
	router.Load(g, middleware.Cors(), middleware.Logger())

	log.Infof("Start to listening the incoming requests on http address: %s", config.GetString("addr"))
	log.Info(http.ListenAndServe(config.GetString("addr"), g).Error())
}
