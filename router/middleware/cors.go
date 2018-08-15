package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"“"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"x-requested-with", "Content-Type", "origin", "authorization", "accept", "client-security-token"},
		AllowCredentials: true,
	})
}
