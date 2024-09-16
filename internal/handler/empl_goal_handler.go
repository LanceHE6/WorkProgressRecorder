package handler

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmplGoalHandler struct {
}

type addEmplGoalRequest struct {
	Status        int    `json:"status" form:"status" binding:"required"`
	TargetCompany string `json:"target_company" form:"target_company" binding:"required"`
	TargetJob     string `json:"target_job" form:"target_job" binding:"required"`
	TargetSalary  int    `json:"target_salary" form:"target_salary" binding:"required"`
	TargetArea    string `json:"target_area" form:"target_area" binding:"required"`
}

// AddEmplGoal
//
//	@Description: 添加就业目标
//	@receiver EmplGoalHandler
//	@param context *gin.Context
func (EmplGoalHandler) AddEmplGoal(context *gin.Context) {
	var data addEmplGoalRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}

	emplGoal := model.EmploymentGoal{
		UID:           userInfo.ID,
		Status:        data.Status,
		TargetCompany: data.TargetCompany,
		TargetJob:     data.TargetJob,
		TargetSalary:  data.TargetSalary,
		TargetArea:    data.TargetArea,
	}
	emplGoalService := service.NewEmplGoalService()
	context.JSON(http.StatusOK, emplGoalService.Insert(emplGoal))
}

// UpdateEmplGoal
//
//	@Description: 更新就业目标
//	@receiver EmplGoalHandler
//	@param context *gin.Context
func (EmplGoalHandler) UpdateEmplGoal(context *gin.Context) {
	var data addEmplGoalRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	emplGoal := model.EmploymentGoal{
		UID:           userInfo.ID,
		Status:        data.Status,
		TargetCompany: data.TargetCompany,
		TargetJob:     data.TargetJob,
		TargetSalary:  data.TargetSalary,
		TargetArea:    data.TargetArea,
	}
	emplGoalService := service.NewEmplGoalService()
	context.JSON(http.StatusOK, emplGoalService.Update(emplGoal))
}
