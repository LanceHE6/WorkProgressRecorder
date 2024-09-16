package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

// EmplGoalRepo
//
//	@Description: 就业目标仓库接口
type EmplGoalRepo interface {
	Insert(emplGoal model.EmploymentGoal) error
	Update(emplGoal model.EmploymentGoal) error
}

// NewEmplGoalRepo
//
//	@Description: 创建就业目标仓库实例
//	@return EmplGoalRepo
func NewEmplGoalRepo() EmplGoalRepo {
	return &impl.EmplGoalRepoImpl{}
}
