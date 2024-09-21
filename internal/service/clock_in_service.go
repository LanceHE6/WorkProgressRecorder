package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

// ClockInService
//
//	@Description: 打卡服务接口
type ClockInService interface {
	AddClockIn(context *gin.Context)
	IsUserClockIn(context *gin.Context)
	SearchClockIn(context *gin.Context)
}

// NewClockInService
//
//	@Description: 创建打卡服务实例
//	@return ClockInService 打卡服务实例
func NewClockInService() ClockInService {
	return &impl.ClockInServiceImpl{}
}
