package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/taskmeter/handler"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		parsed, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set("uid", parsed.ID)
		c.Set("username", parsed.Username)
		c.Next()
	}
}
