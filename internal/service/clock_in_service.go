package service

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

// ClockInService
//
//	@Description: 打卡服务接口
type ClockInService interface {
	AddClockIn(clockInData model.ClockIn) *pkg.Response
}

// NewClockInService
//
//	@Description: 创建打卡服务实例
//	@return ClockInService 打卡服务实例
func NewClockInService() ClockInService {
	return &impl.ClockInServiceImpl{}
}
