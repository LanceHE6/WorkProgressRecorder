package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
	"WorkProgressRecord/pkg"
)

// ClockInRepo
//
//	@Description: 打卡记录仓库接口
type ClockInRepo interface {
	Insert(clockInData model.ClockIn) error
	SelectUserLatest(uid int64) (model.ClockIn, error)
	SearchClockIns(params pkg.SearchClockInParams) ([]model.ClockIn, int, error)
}

// NewClockInRepo
//
//	@Description: 创建打卡记录仓库实例
//	@return ClockInRepo 打卡记录仓库实例
func NewClockInRepo() ClockInRepo {
	return &impl.ClockInRepoImpl{}
}
