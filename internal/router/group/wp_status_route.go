package group

import (
	"WorkProgressRecord/internal/handler"
	"github.com/gin-gonic/gin"
)

// WpStatusRoute
//
//	@Description: 就业状态路由
//	@param group 路由组
func WpStatusRoute(group *gin.RouterGroup) {
	wPStatusHandler := handler.WpStatusHandler{}

	wPStatusRoute := group.Group("/wps")
	wPStatusRoute.GET("/list", func(context *gin.Context) {
		wPStatusHandler.GetStatus(context)
	})
}
