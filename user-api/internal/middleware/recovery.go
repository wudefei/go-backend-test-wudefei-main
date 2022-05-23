package middleware

import (
	"io"
	"net/http"

	"user-api/internal/util"

	"github.com/gin-gonic/gin"
)

// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter(gin.DefaultErrorWriter)
}

// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func RecoveryWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//友好打印异常
				logID := c.GetString("logId")
				util.PrettyPanic(logID, "中间件捕抓到异常", r)

				//返回报错信息
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    util.InternalServerError,
					"message": util.ErrorMsg[util.InternalServerError],
					"data":    map[string]string{},
				})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
