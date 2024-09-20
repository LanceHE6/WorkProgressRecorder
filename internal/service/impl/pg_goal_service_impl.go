package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PGGoalServiceImpl
//
//	@Description: 考研目标服务实现
type PGGoalServiceImpl struct {
}

// InsertAndUpdate
//
//	@Description: 添加或更新考研目标
//	@receiver e PGGoalServiceImpl
//	@param emplGoal model.PostgraduateGoal
//	@return *pkg.Response 返回响应
func (p PGGoalServiceImpl) InsertAndUpdate(context *gin.Context) {
	type addPGGoalRequest struct {
		TargetUniversity string  `json:"target_university" form:"target_university" binding:"required"`
		TargetMajor      string  `json:"target_major" form:"target_university" binding:"required"`
		TargetScore      float64 `json:"target_score" form:"target_university" binding:"required"`
	}
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
	pgGoalRepo := repo.NewPGGoalRepo()
	oldPgGoal, err := pgGoalRepo.SelectByUID(pgGoal.UID)
	if err != nil || oldPgGoal.ID == 0 {
		err = pgGoalRepo.Insert(pgGoal)
		if err != nil {
			context.JSON(http.StatusOK, pkg.ErrorResponse(1, "添加失败", err))
			return
		}
	} else {
		pgGoal.ID = oldPgGoal.ID
		err = pgGoalRepo.Update(pgGoal)
		if err != nil {
			context.JSON(http.StatusOK, pkg.ErrorResponse(1, "更新失败", err))
			return
		}
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}
