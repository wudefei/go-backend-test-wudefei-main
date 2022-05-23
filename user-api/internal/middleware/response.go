package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"user-api/internal/data"
	"user-api/internal/util"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

// Write 响应
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

var simpleLogPath []string
var notLogPath []string

// SetPathSimpleLog 设置简单日志的path
func SetPathSimpleLog(pathes []string) {
	simpleLogPath = pathes
}

// SetNotLogPath 不打印日志的path
func SetNotLogPath(pathes []string) {
	notLogPath = pathes
}

// Response 打印请求 和 响应
func Response(c *gin.Context) {
	var (
		logID = util.Uniqid()
	)
	c.Set("logId", logID)

	// 把request的内容读取出来,刚刚读出来的要再写进去
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	//重置writer，需要打印resp
	blw := bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	//请求打日志
	if !util.ContainString(notLogPath, c.Request.RequestURI) {
		if util.ContainString(simpleLogPath, c.Request.RequestURI) {
			fmt.Println(logID, "request", c.Request.RequestURI, c.Request.Method)
		} else {
			fmt.Println(logID, "request", c.Request.RequestURI, c.Request.Method, strings.Replace(string(bodyBytes), "\n", "", -1))
		}
	}

	//业务
	c.Next()

	//响应打日志
	respLogArray := []interface{}{logID, "response", c.Request.RequestURI, " ", c.Request.Method, c.Writer.Status()}
	if !util.ContainString(simpleLogPath, c.Request.RequestURI) {
		respLogArray = append(respLogArray, blw.bodyBuf.String())
	}

	if !util.ContainString(notLogPath, c.Request.RequestURI) {
		fmt.Println(respLogArray...)
	}

	//response已有，返回
	if c.Writer.Written() {
		return
	}

	rp := data.ResponseBase{
		Code:    util.InternalServerError,
		Message: "service error",
		Data:    nil,
	}
	c.JSON(http.StatusInternalServerError, rp)

	c.Abort()

}
