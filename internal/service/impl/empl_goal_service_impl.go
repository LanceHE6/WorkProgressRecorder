package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EmplGoalServiceImpl
//
//	@Description: 就业目标服务实现
type EmplGoalServiceImpl struct {
}

// InsertAndUpdate
//
//	@Description: 添加或更新就业目标
//	@receiver e EmplGoalServiceImpl
//	@param emplGoal model.EmploymentGoal
//	@return *pkg.Response 返回响应
func (e EmplGoalServiceImpl) InsertAndUpdate(context *gin.Context) {
	type addEmplGoalRequest struct {
		Status        int    `json:"status" form:"status" binding:"required"`
		TargetCompany string `json:"target_company" form:"target_company" binding:"required"`
		TargetJob     string `json:"target_job" form:"target_job" binding:"required"`
		TargetSalary  string `json:"target_salary" form:"target_salary" binding:"required"`
		TargetArea    string `json:"target_area" form:"target_area" binding:"required"`
	}
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
	emplGoalRepo := repo.NewEmplGoalRepo()
	oldEmplGoal, err := emplGoalRepo.SelectByUID(emplGoal.UID)
	// 不存在则添加
	if err != nil || oldEmplGoal.ID == 0 {
		err = emplGoalRepo.Insert(emplGoal)
		if err != nil {
			context.JSON(http.StatusOK, pkg.ErrorResponse(1, "添加失败", err))
			return
		}
	} else {
		// 存在则更新
		emplGoal.ID = oldEmplGoal.ID
		err = emplGoalRepo.Update(emplGoal)
		if err != nil {
			context.JSON(http.StatusOK, pkg.ErrorResponse(1, "更新失败", err))
			return
		}
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}
