package handler

import (
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// WpStatusHandler
//
//	@Description: 就业状态处理器
type WpStatusHandler struct {
}

// GetStatus
//
//	@Description: 获取状态列表
//	@receiver WpStatusHandler
//	@param context
func (WpStatusHandler) GetStatus(context *gin.Context) {
	context.JSON(200, service.NewWpStatusService().GetStatus())
}
