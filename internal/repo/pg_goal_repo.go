package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

// PGGoalRepo
//
//	@Description: 考研目标仓库接口
type PGGoalRepo interface {
	Insert(pgGoal model.PostgraduateGoal) error
	Update(pgGoal model.PostgraduateGoal) error
}

// NewPGGoalRepo
//
//	@Description: 创建考研目标仓库实例
//	@return EmplGoalRepo
func NewPGGoalRepo() PGGoalRepo {
	return &impl.PGGoalRepoImpl{}
}
