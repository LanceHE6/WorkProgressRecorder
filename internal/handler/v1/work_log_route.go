package v1

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// WorkLogRoute
//
//	@Description: 工作日志路由
//	@param group *gin.RouterGroup
func WorkLogRoute(group *gin.RouterGroup) {
	workLogService := service.NewWorkLogService()

	workLogGroup := group.Group("/wl")
	workLogGroup.Use(middleware.AuthMiddleware())
	workLogGroup.POST("/add", func(context *gin.Context) {
		workLogService.AddWorkLog(context)
	})
	workLogGroup.PUT("/status/add", func(context *gin.Context) {
		workLogService.AddStatusTimeLine(context)
	})
	workLogGroup.GET("/list", func(context *gin.Context) { // 获取工作日志列表
		workLogService.GetUserWorkLog(context)
	})
}
