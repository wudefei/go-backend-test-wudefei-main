package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"user-api/internal/data"
	"user-api/internal/db/redis"
	"user-api/internal/util"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "not allowed"})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(500, gin.H{"message": "cannot find the handler"})
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		tokenInput := c.GetHeader("token")
		phone := c.Query("phone")
		if tokenInput == "" {
			fmt.Println("Invalid request")
			c.JSON(http.StatusOK, data.ResponseBase{
				Code:    util.Invalid,
				Message: "Invalid request",
			})
			return
		}
		logID := util.Uniqid()

		//检查token是否过期
		token, err := redis.NewRedisHelper(logID).StringGet(data.TokenPrefix + phone)
		if err != nil || token == "" {
			c.JSON(http.StatusOK, data.ResponseBase{
				Code:    util.InternalServerError,
				Message: "internal error",
			})
			return
		}

		//校验token是否一致
		if token != tokenInput {
			c.JSON(http.StatusOK, data.ResponseBase{
				Code:    util.Invalid,
				Message: "Invalid request",
			})
			return
		}

		fmt.Println("auth success")
		c.Next()
	}
}

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
