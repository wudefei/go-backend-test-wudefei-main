package service

import (
	"user-api/internal/data"

	"github.com/gin-gonic/gin"
)

// common service
type CommonService struct {
	ctx   *gin.Context
	logID string
}

func NewCommonService(ctx *gin.Context, logId string) *CommonService {
	return &CommonService{
		ctx:   ctx,
		logID: logId,
	}
}

func (c *CommonService) HealthCheck(req data.HealthCheckReq) (string, error) {
	return "ok", nil
}
