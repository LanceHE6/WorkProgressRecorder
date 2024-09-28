package v1

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// EmplGoalRoute
//
//	@Description: 就业目标路由组
//	@param group *gin.RouterGroup 路由组
func EmplGoalRoute(group *gin.RouterGroup) {
	emplGoalService := service.NewEmplGoalService()

	emplGoalGroup := group.Group("/emgl")
	emplGoalGroup.Use(middleware.AuthMiddleware())
	emplGoalGroup.POST("/add", func(context *gin.Context) {
		emplGoalService.InsertAndUpdate(context)
	})

}
