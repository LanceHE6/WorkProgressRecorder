package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkLogServiceImpl struct {
}

// AddWorkLog
//
//	@Description: 添加工作日志
//	@receiver w WorkLogServiceImpl
//	@param context *gin.Context
func (w WorkLogServiceImpl) AddWorkLog(context *gin.Context) {
	type AddWorkLogReq struct {
		CompanyName string `json:"company_name" form:"company_name" binding:"required"`
		Job         string `json:"job" form:"job" binding:"required"`
		Salary      string `json:"salary" form:"salary" binding:"required"`
		Location    string `json:"location" form:"location" binding:"required"`
	}
	var data AddWorkLogReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
		return
	}
	workLog := model.WorkLog{
		UID:            userInfo.ID,
		CompanyName:    data.CompanyName,
		Job:            data.Job,
		Salary:         data.Salary,
		Location:       data.Location,
		StatusTimeLine: make(model.StatusTimeLine, 0),
	}
	workLogRepo := repo.NewWorkLogRepo()
	if err := workLogRepo.InsertWorkLog(workLog); err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(1, "创建新工作日志失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// AddStatusTimeLine
//
//	@Description: 添加状态时间线
//	@receiver w WorkLogServiceImpl
//	@param context *gin.Context
func (w WorkLogServiceImpl) AddStatusTimeLine(context *gin.Context) {
	type AddStatusTimeLineReq struct {
		ID         int64  `json:"id" form:"id" binding:"required"`
		CreateTime int64  `json:"create_time" form:"create_time" binding:"required"`
		Status     string `json:"status" form:"status" binding:"required"`
	}
	var data AddStatusTimeLineReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	statusTime := model.StatusTime{
		CreatedAt: data.CreateTime,
		Status:    data.Status,
	}
	workLogRepo := repo.NewWorkLogRepo()
	err := workLogRepo.AddStatusTimeLine(data.ID, statusTime)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(1, "添加状态时间线失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// GetUserWorkLog
//
//	@Description: 获取用户工作日志
//	@receiver w WorkLogServiceImpl
//	@param context *gin.Context
func (w WorkLogServiceImpl) GetUserWorkLog(context *gin.Context) {
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
		return
	}
	workLogRepo := repo.NewWorkLogRepo()
	workLogs, err := workLogRepo.SelectByUID(userInfo.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(1, "获取用户工作日志失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"work_logs": workLogs,
	}))
}
