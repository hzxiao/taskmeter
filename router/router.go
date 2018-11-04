package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/taskmeter/handler"
	"github.com/hzxiao/taskmeter/router/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler.Register(g)
	return g
}
