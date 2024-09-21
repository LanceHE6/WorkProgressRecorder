package handler

import (
	"WorkProgressRecord/internal/handler/group"
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
	group.PGGoalRoute(apiV1)
	group.ClockInRoute(apiV1)
	group.PgCountDownRoute(apiV1)
	group.WorkLogRoute(apiV1)
	group.PostRoute(apiV1)
}
