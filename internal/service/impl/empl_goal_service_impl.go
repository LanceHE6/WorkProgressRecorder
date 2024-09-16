package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

type EmplGoalServiceImpl struct {
}

func (e EmplGoalServiceImpl) Insert(emplGoal model.EmploymentGoal) *pkg.Response {
	emplGoalRepo := repo.NewEmplGoalRepo()
	err := emplGoalRepo.Insert(emplGoal)
	if err != nil {
		return pkg.NewResponse(100, "添加失败", err.Error())
	}
	return pkg.SuccessResponse(nil)
}
