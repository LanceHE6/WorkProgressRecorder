package handler

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ClockInHandler struct {
}

type AddClockInRequest struct {
	Location string `json:"location" form:"location" binding:"required"`
}

// AddClockIn
//
//	@Description: 打卡
//	@receiver ClockInHandler 打卡
//	@param context 上下文
func (ClockInHandler) AddClockIn(context *gin.Context) {
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
	clockInService := service.NewClockInService()
	context.JSON(http.StatusOK, clockInService.AddClockIn(clockInData))
}
