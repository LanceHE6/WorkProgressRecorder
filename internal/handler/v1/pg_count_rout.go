package v1

import (
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// PgCountDownRoute
//
//	@Description: 考研倒计时路由
//	@param group *gin.RouterGroup 路由组
func PgCountDownRoute(group *gin.RouterGroup) {
	pgCDService := service.NewPgCountdownService()
	pgGoalGroup := group.Group("/pgee")
	pgGoalGroup.GET("/cd", func(context *gin.Context) {
		pgCDService.GetPgCountdown(context)
	})
}
