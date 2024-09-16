package service

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

// EmplGoalService
//
//	@Description: 就业目标服务接口
type EmplGoalService interface {
	Insert(emplGoal model.EmploymentGoal) *pkg.Response
	Update(emplGoal model.EmploymentGoal) *pkg.Response
}

// NewEmplGoalService
//
//	@Description: 创建就业目标服务实例
//	@return EmplGoalService 就业目标服务实例
func NewEmplGoalService() EmplGoalService {
	return &impl.EmplGoalServiceImpl{}
}
