package group

import (
	"WorkProgressRecord/internal/handler"
	"WorkProgressRecord/internal/middleware"
	"github.com/gin-gonic/gin"
)

// PGGoalRoute
//
//	@Description: 考研目标路由组
//	@param group *gin.RouterGroup 路由组
func PGGoalRoute(group *gin.RouterGroup) {
	pgGoalHandler := handler.PGGoalHandler{}

	pgGoalGroup := group.Group("/pggl")
	pgGoalGroup.Use(middleware.AuthMiddleware())
	pgGoalGroup.POST("/add", func(context *gin.Context) {
		pgGoalHandler.AddPGGoal(context)
	})
	pgGoalGroup.PUT("/update", func(context *gin.Context) {
		pgGoalHandler.UpdatePGGoal(context)
	})
}
