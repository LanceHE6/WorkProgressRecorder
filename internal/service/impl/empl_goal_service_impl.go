package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// EmplGoalServiceImpl
//
//	@Description: 就业目标服务实现
type EmplGoalServiceImpl struct {
}

// Insert
//
//	@Description: 添加就业目标
//	@receiver e EmplGoalServiceImpl
//	@param emplGoal model.EmploymentGoal
//	@return *pkg.Response 返回响应
func (e EmplGoalServiceImpl) Insert(emplGoal model.EmploymentGoal) *pkg.Response {
	emplGoalRepo := repo.NewEmplGoalRepo()
	err := emplGoalRepo.Insert(emplGoal)
	if err != nil {
		return pkg.ErrorResponse(1, "添加失败", err)
	}
	return pkg.SuccessResponse(nil)
}

// Update
//
//	@Description: 更新就业目标
//	@receiver e EmplGoalServiceImpl
//	@param emplGoal model.EmploymentGoal
//	@return *pkg.Response 返回响应
func (e EmplGoalServiceImpl) Update(emplGoal model.EmploymentGoal) *pkg.Response {
	emplGoalRepo := repo.NewEmplGoalRepo()
	err := emplGoalRepo.Update(emplGoal)
	if err != nil {
		return pkg.ErrorResponse(1, "更新失败", err)
	}
	return pkg.SuccessResponse(nil)
}
