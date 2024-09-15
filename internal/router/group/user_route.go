package group

import (
	"WorkProgressRecord/internal/handler"
	"github.com/gin-gonic/gin"
)

// UserRoute
//
//	@Description: 用户路由组
//	@param group *gin.RouterGroup 路由组
func UserRoute(group *gin.RouterGroup) {
	userHandler := handler.UserHandler{}

	userGroup := group.Group("/user")
	userGroup.POST("/login", func(context *gin.Context) {
		userHandler.Login(context)
	})
}
