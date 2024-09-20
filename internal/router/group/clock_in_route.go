package group

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// ClockInRoute
//
//	@Description: 打卡路由
//	@param group gin.RouterGroup
func ClockInRoute(group *gin.RouterGroup) {
	clockInService := service.NewClockInService()

	clockInGroup := group.Group("/clock_in")
	clockInGroup.Use(middleware.AuthMiddleware())
	clockInGroup.POST("/add", func(context *gin.Context) {
		clockInService.AddClockIn(context)
	})
}
