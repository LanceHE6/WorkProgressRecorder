package handler

import (
	"WorkProgressRecord/internal/handler/v1"
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

	v1.UserRoute(apiV1)
	v1.EmplGoalRoute(apiV1)
	v1.PGGoalRoute(apiV1)
	v1.ClockInRoute(apiV1)
	v1.PgCountDownRoute(apiV1)
	v1.WorkLogRoute(apiV1)
	v1.PostRoute(apiV1)
	v1.VersionRoute(apiV1)
}
