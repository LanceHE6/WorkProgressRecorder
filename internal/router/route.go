package router

import (
	"WorkProgressRecord/internal/router/group"
	"github.com/gin-gonic/gin"
)

// Route
//
//	@Description: 总路由
//	@param gin *gin.Engine gin框架
func Route(gin *gin.Engine) {
	api := gin.Group("/api")
	apiV1 := api.Group("/v1")

	// 用户路由组

	group.UserRoute(apiV1)
	group.EmplGoalRoute(apiV1)
}
