package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

// PgCountdownService
//
//	@Description: 倒计时服务
type PgCountdownService interface {
	GetPgCountdown(context *gin.Context)
}

// NewPgCountdownService
//
//	@Description: 创建倒计时服务
//	@return PgCountdownService
func NewPgCountdownService() PgCountdownService {
	return &impl.PgCountdownServiceImpl{}
}
