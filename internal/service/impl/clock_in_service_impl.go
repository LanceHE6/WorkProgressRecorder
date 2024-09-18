package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// ClockInServiceImpl
//
//	@Description: 打卡服务实现
type ClockInServiceImpl struct {
}

// AddClockIn
//
//	@Description: 添加打卡记录
//	@receiver ClockInServiceImpl 打卡服务
//	@param clockInData 打卡数据
//	@return *pkg.Response 响应
func (ClockInServiceImpl) AddClockIn(clockInData model.ClockIn) *pkg.Response {
	clockInRepo := repo.NewClockInRepo()
	err := clockInRepo.Insert(clockInData)
	if err != nil {
		return pkg.ErrorResponse(1, "打卡失败", err)
	}
	return pkg.SuccessResponse(nil)
}
