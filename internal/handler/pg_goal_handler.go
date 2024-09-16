package handler

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PGGoalHandler struct {
}

type addPGGoalRequest struct {
	TargetUniversity string  `json:"target_university" form:"target_university" binding:"required"`
	TargetMajor      string  `json:"target_major" form:"target_university" binding:"required"`
	TargetScore      float64 `json:"target_score" form:"target_university" binding:"required"`
}

// AddPGGoal
//
//	@Description: 添加考研目标
//	@receiver PGGoalHandler
//	@param context *gin.Context
func (PGGoalHandler) AddPGGoal(context *gin.Context) {
	var data addPGGoalRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}

	pgGoal := model.PostgraduateGoal{
		UID:              userInfo.ID,
		TargetUniversity: data.TargetUniversity,
		TargetMajor:      data.TargetMajor,
		TargetScore:      data.TargetScore,
	}
	pgGoalService := service.NewPGGoalService()
	context.JSON(http.StatusOK, pgGoalService.Insert(pgGoal))
}

// UpdatePGGoal
//
//	@Description: 更新考研目标
//	@receiver PGGoalHandler
//	@param context *gin.Context
func (PGGoalHandler) UpdatePGGoal(context *gin.Context) {
	var data addPGGoalRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	pgGoal := model.PostgraduateGoal{
		UID:              userInfo.ID,
		TargetUniversity: data.TargetUniversity,
		TargetMajor:      data.TargetMajor,
		TargetScore:      data.TargetScore,
	}
	pgGoalService := service.NewPGGoalService()
	context.JSON(http.StatusOK, pgGoalService.Update(pgGoal))
}
