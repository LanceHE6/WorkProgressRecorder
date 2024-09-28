package v1

import (
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// VersionRoute
//
//	@Description: 版本路由
//	@param group *gin.RouterGroup
func VersionRoute(group *gin.RouterGroup) {
	versionService := service.NewVersionService()

	group.GET("/ver", func(context *gin.Context) {
		versionService.GetVersion(context)
	})
	group.GET("/ping", func(context *gin.Context) {
		versionService.GetVersion(context)
	})
}
