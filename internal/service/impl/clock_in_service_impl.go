package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	// 判断用户当前时间段是否已经打卡
	clockInRepo := repo.NewClockInRepo()
	latestClockIn, err := clockInRepo.SelectUserLatest(userInfo.ID)
	if err != nil {
		if err.Error() != "record not found" {
			context.JSON(http.StatusOK, pkg.ErrorResponse(-2, "打卡失败", err))
			return
		}
	}
	if pkg.IsInCurrentPeriod(latestClockIn.ClockInTime) {
		context.JSON(http.StatusOK, pkg.FailedResponse(1, "您在当前时间段已经打过卡了哦"))
		return
	}
	clockInData := model.ClockIn{
		UID:             userInfo.ID,
		ClockInLocation: data.Location,
		ClockInTime:     timestamp,
	}

	err = clockInRepo.Insert(clockInData)
	if err != nil {
		context.JSON(http.StatusOK, pkg.ErrorResponse(1, "打卡失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// IsUserClockIn
//
//	@Description: 判断用户当前时间段是否已经打卡
//	@receiver i ClockInServiceImpl
//	@param context gin.Context
func (i ClockInServiceImpl) IsUserClockIn(context *gin.Context) {
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	// 判断用户当前时间段是否已经打卡
	clockInRepo := repo.NewClockInRepo()
	latestClockIn, err := clockInRepo.SelectUserLatest(userInfo.ID)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusOK, pkg.NewResponse(0, "您还未打卡哦", nil))
			return
		}
		context.JSON(http.StatusOK, pkg.ErrorResponse(-2, "未知错误", err))
		return
	}
	if pkg.IsInCurrentPeriod(latestClockIn.ClockInTime) {
		context.JSON(http.StatusOK, pkg.FailedResponse(1, "您在当前时间段已经打过卡了哦"))
		return
	} else {
		context.JSON(http.StatusOK, pkg.NewResponse(0, "您还未打卡哦", nil))
	}
}

func (i ClockInServiceImpl) SearchClockIn(context *gin.Context) {
	var params pkg.SearchClockInParams
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	params.Page = &page
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	params.Limit = &limit
	if context.Query("uid") != "" {
		uid, _ := strconv.Atoi(context.Query("uid"))
		uid64 := int64(uid)
		params.UID = &uid64
	}
	if context.Query("time_slot") != "" {
		timeSlot := context.Query("time_slot")
		params.TimeSlot = (*pkg.TimeSlot)(&timeSlot)
	}
	if context.Query("date") != "" {
		date := context.Query("date")
		params.Date = &date
	}
	clockInRepo := repo.NewClockInRepo()
	clockInList, count, err := clockInRepo.SearchClockIns(params)
	if err != nil {
		context.JSON(http.StatusOK, pkg.ErrorResponse(-1, "查询失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"rows":      clockInList,
		"total":     count,
		"page":      page,
		"page_size": limit,
	}))
}
