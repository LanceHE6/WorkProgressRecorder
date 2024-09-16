package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

type EmplGoalRepo interface {
	Insert(emplGoal model.EmploymentGoal) error
}

func NewEmplGoalRepo() EmplGoalRepo {
	return &impl.EmplGoalRepoImpl{}
}
