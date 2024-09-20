package service

import (
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

// PgCountdownService
//
//	@Description: 倒计时服务
type PgCountdownService interface {
	GetPgCountdown() *pkg.Response
}

// NewPgCountdownService
//
//	@Description: 创建倒计时服务
//	@return PgCountdownService
func NewPgCountdownService() PgCountdownService {
	return &impl.PgCountdownServiceImpl{}
}
