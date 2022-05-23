package imp

import (
	"net/http"

	"user-api/internal/data"
	"user-api/internal/service"

	"user-api/internal/util"

	"github.com/gin-gonic/gin"
)

/**
 * 通用的服务接口,例如健康检查
 *
 */
type CommonApiImp struct{}

func NewCommonApiImp() *CommonApiImp {
	return &CommonApiImp{}
}

func (common *CommonApiImp) HealthCheck(c *gin.Context) {
	var (
		logID   = util.Uniqid()
		message = "success"
	)
	healthData, err := service.NewCommonService(c, logID).HealthCheck(data.HealthCheckReq{})
	if err != nil {
		message = err.Error()
	}
	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: message,
		Data:    healthData,
	})
}
