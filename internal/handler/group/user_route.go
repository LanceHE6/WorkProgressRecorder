package group

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// UserRoute
//
//	@Description: 用户路由组
//	@param group *gin.RouterGroup 路由组
func UserRoute(group *gin.RouterGroup) {
	userService := service.NewUserService()

	userGroup := group.Group("/user")
	userGroup.POST("/login", func(context *gin.Context) {
		userService.Login(context)
	})
	userGroup.POST("/import", middleware.IsAdminMiddleware(), func(context *gin.Context) {
		userService.Import(context)
	})
	userGroup.PUT("/update/psw", middleware.AuthMiddleware(), func(context *gin.Context) {
		userService.UpdatePassword(context)
	})
	userGroup.PUT("/update/info", middleware.AuthMiddleware(), func(context *gin.Context) {
		userService.UpdateUserInfo(context)
	})
	userGroup.GET("/target", middleware.AuthMiddleware(), func(context *gin.Context) {
		userService.GetUserTargetInfo(context)
	})
	userGroup.GET("/info", middleware.AuthMiddleware(), func(context *gin.Context) {
		userService.GetUserInfo(context)
	})
	userGroup.GET("/search", middleware.IsAdminMiddleware(), func(context *gin.Context) {
		userService.SearchUsers(context)
	})
	userGroup.DELETE("/dire/del", middleware.AuthMiddleware(), func(context *gin.Context) {
		userService.DeleteUserDirection(context)
	})
}
