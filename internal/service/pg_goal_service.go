package service

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

// PGGoalService
//
//	@Description: 考研目标服务接口
type PGGoalService interface {
	Insert(pgGoal model.PostgraduateGoal) *pkg.Response
	Update(pgGoal model.PostgraduateGoal) *pkg.Response
}

// NewPGGoalService
//
//	@Description: 创建考研目标服务实例
//	@return PGGoalService 考研目标服务实例
func NewPGGoalService() PGGoalService {
	return &impl.PGGoalServiceImpl{}
}
