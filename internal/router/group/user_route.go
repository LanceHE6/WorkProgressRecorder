package group

import (
	"WorkProgressRecord/internal/handler"
	"WorkProgressRecord/internal/middleware"
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
	userGroup.POST("/import", middleware.AuthMiddleware(), func(context *gin.Context) {
		userHandler.ImportUser(context)
	})
	userGroup.PUT("/update/psw", middleware.AuthMiddleware(), func(context *gin.Context) {
		userHandler.UpdatePassword(context)
	})
	userGroup.GET("/target", middleware.AuthMiddleware(), func(context *gin.Context) {
		userHandler.GetTargetInfo(context)
	})
}
