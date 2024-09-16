package service

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

type EmplGoalService interface {
	Insert(emplGoal model.EmploymentGoal) *pkg.Response
}

func NewEmplGoalService() EmplGoalService {
	return &impl.EmplGoalServiceImpl{}
}
