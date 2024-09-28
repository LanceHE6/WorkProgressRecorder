package v1

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// PGGoalRoute
//
//	@Description: 考研目标路由组
//	@param group *gin.RouterGroup 路由组
func PGGoalRoute(group *gin.RouterGroup) {
	pgGoalService := service.NewPGGoalService()

	pgGoalGroup := group.Group("/pggl")
	pgGoalGroup.Use(middleware.AuthMiddleware())
	pgGoalGroup.POST("/add", func(context *gin.Context) {
		pgGoalService.InsertAndUpdate(context)
	})
}
