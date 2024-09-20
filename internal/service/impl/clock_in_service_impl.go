package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ClockInServiceImpl
//
//	@Description: 打卡服务实现
type ClockInServiceImpl struct {
}

// AddClockIn
//
//	@Description: 添加打卡记录
//	@receiver ClockInServiceImpl 打卡服务
//	@param clockInData 打卡数据
//	@return *pkg.Response 响应
func (ClockInServiceImpl) AddClockIn(context *gin.Context) {
	type AddClockInRequest struct {
		Location string `json:"location" form:"location" binding:"required"`
	}
	var data AddClockInRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	clockInData := model.ClockIn{
		UID:             userInfo.ID,
		ClockInLocation: data.Location,
		ClockInTime:     timestamp,
	}

	clockInRepo := repo.NewClockInRepo()
	err = clockInRepo.Insert(clockInData)
	if err != nil {
		context.JSON(http.StatusOK, pkg.ErrorResponse(1, "打卡失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}
