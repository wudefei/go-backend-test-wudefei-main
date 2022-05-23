package router

import (
	"net/http"
	"time"

	"user-api/internal/imp"
	mdw "user-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 路由
func InitRouter() *gin.Engine {
	router := gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "interface not exist",
			"data":    map[string]string{},
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "request is not allowed",
			"data":    map[string]string{},
		})
	})

	router.Any("/", func(c *gin.Context) {
		c.Header("server", "1.0.0")
		c.JSON(200, gin.H{
			"message": "handle service " + time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	//不打印日志的path
	mdw.SetNotLogPath([]string{
		"/healthz",
	})

	//使用中间件捕捉异常
	router.Use(mdw.Recovery())

	//响应
	router.Use(mdw.Response)

	//健康检查
	commonServer := imp.NewCommonApiImp()
	commonApi := router.Group("")
	{
		commonApi.Any("/healthz", commonServer.HealthCheck) //healthy check
	}

	//业务接口
	user := imp.NewUserApiImp()
	interApi := router.Group("/v1")
	{
		//跳过用户权限校验检查的接口
		var authUrlArr []string
		authUrlArr = append(authUrlArr, "/healthz", "/sms", "/login")
		interApi.Use(mdw.UserAuthMiddleware(
			mdw.AllowPathPrefixSkipper(authUrlArr...),
		))

		interApi.POST("/sms", user.SendSms)
		interApi.POST("/login", user.UserLogin)
		interApi.GET("/user", user.QueryUser)
		interApi.POST("/user", user.AddUser)
		interApi.PUT("/user", user.UpdateUser)
		interApi.DELETE("/user", user.DeleteUser)

	}

	return router
}
