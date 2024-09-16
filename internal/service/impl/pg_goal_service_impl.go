package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// PGGoalServiceImpl
//
//	@Description: 考研目标服务实现
type PGGoalServiceImpl struct {
}

// Insert
//
//	@Description: 添加考研目标
//	@receiver e PGGoalServiceImpl
//	@param emplGoal model.PostgraduateGoal
//	@return *pkg.Response 返回响应
func (p PGGoalServiceImpl) Insert(pgGoal model.PostgraduateGoal) *pkg.Response {
	emplGoalRepo := repo.NewPGGoalRepo()
	err := emplGoalRepo.Insert(pgGoal)
	if err != nil {
		return pkg.ErrorResponse(1, "添加失败", err)
	}
	return pkg.SuccessResponse(nil)
}

// Update
//
//	@Description: 更新考研目标
//	@receiver e PGGoalServiceImpl
//	@param emplGoal model.PostgraduateGoal
//	@return *pkg.Response 返回响应
func (p PGGoalServiceImpl) Update(pgGoal model.PostgraduateGoal) *pkg.Response {
	emplGoalRepo := repo.NewPGGoalRepo()
	err := emplGoalRepo.Update(pgGoal)
	if err != nil {
		return pkg.ErrorResponse(1, "更新失败", err)
	}
	return pkg.SuccessResponse(nil)
}
