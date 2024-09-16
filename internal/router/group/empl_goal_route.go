package group

import (
	"WorkProgressRecord/internal/handler"
	"WorkProgressRecord/internal/middleware"
	"github.com/gin-gonic/gin"
)

// EmplGoalRoute
//
//	@Description: 就业目标路由组
//	@param group *gin.RouterGroup 路由组
func EmplGoalRoute(group *gin.RouterGroup) {
	emplGoalHandler := handler.EmplGoalHandler{}

	emplGoalGroup := group.Group("/emgl")
	emplGoalGroup.Use(middleware.AuthMiddleware())
	emplGoalGroup.POST("/add", func(context *gin.Context) {
		emplGoalHandler.AddEmplGoal(context)
	})
}
