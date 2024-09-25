package group

import (
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

func VersionRoute(group *gin.RouterGroup) {
	versionService := service.NewVersionService()

	group.GET("/ver", func(context *gin.Context) {
		versionService.GetVersion(context)
	})
	group.GET("/ping", func(context *gin.Context) {
		versionService.GetVersion(context)
	})
}
