package data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	VerifiCodePrefix string = "verificode:"
	TokenPrefix      string = "token:"
	Token            string = "token"
)

// 响应中的通用字段
type ResponseBase struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 健康检查请求体
type HealthCheckReq struct{}

// 响应
func ResCode(c *gin.Context, msg string, code int32) {
	ret := ResponseBase{Code: code, Message: msg}
	ResJSON(c, &ret)
}

// 响应JSON数据
func ResJSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
	c.Abort()
}
