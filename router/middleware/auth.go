package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/config"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/hzxiao/taskmeter/pkg/token"
	"github.com/lexkong/log"
	"net/http"
	"strconv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		parsed, err := token.ParseRequest(c)
		if err != nil {
			log.Error("[Auth]", err)
			c.JSON(http.StatusOK, goutil.Map{
				"code":    errno.ErrUnauthorized.Code,
				"message": errno.ErrUnauthorized.Message,
			})
			c.Abort()
			return
		}

		generate, _ := strconv.ParseInt(parsed.GenerateTime, 10, 64)
		validity := config.GetInt64("token_validity")
		if validity > 0 && generate+validity < timeutil.Now() {
			log.Infof("[Auth] uid(%v) invalid token", parsed.ID)
			c.JSON(http.StatusOK, goutil.Map{
				"code":    errno.ErrUnauthorized.Code,
				"message": errno.ErrUnauthorized.Message,
			})
			c.Abort()
			return
		}
		c.Set("uid", parsed.ID)
		c.Set("username", parsed.Username)
		c.Set("source", parsed.Source)
		c.Next()
	}
}
