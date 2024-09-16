package handler

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmplGoalHandler struct {
}

type insertRequest struct {
	UID           int64  `json:"uid" form:"uid" binding:"required"`
	Status        int    `json:"status" form:"status" binding:"required"`
	TargetCompany string `json:"target_company" form:"target_company" binding:"required"`
	TargetJob     string `json:"target_job" form:"target_job" binding:"required"`
	TargetSalary  int    `json:"target_salary" form:"target_salary" binding:"required"`
	TargetArea    string `json:"target_area" form:"target_area" binding:"required"`
}

func (EmplGoalHandler) AddEmplGoal(context *gin.Context) {
	var data insertRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	emplGoal := model.EmploymentGoal{
		UID:           data.UID,
		Status:        data.Status,
		TargetCompany: data.TargetCompany,
		TargetJob:     data.TargetJob,
		TargetSalary:  data.TargetSalary,
		TargetArea:    data.TargetArea,
	}
	emplGoalService := service.NewEmplGoalService()
	context.JSON(http.StatusOK, emplGoalService.Insert(emplGoal))
}
