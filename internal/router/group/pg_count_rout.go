package group

import (
	"WorkProgressRecord/internal/handler"
	"github.com/gin-gonic/gin"
)

// PgCountDownRoute
//
//	@Description: 考研倒计时路由
//	@param group *gin.RouterGroup 路由组
func PgCountDownRoute(group *gin.RouterGroup) {
	pgCDHandler := handler.PgCountdownHandler{}
	pgGoalGroup := group.Group("/pgee")
	pgGoalGroup.GET("/cd", func(context *gin.Context) {
		pgCDHandler.GetPgCountdown(context)
	})
}
