package impl

import (
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
)

type VersionServiceImpl struct {
}

// GetVersion
//
//	@Description: 获取版本信息
//	@receiver v VersionServiceImpl
//	@param context *gin.Context
func (v VersionServiceImpl) GetVersion(context *gin.Context) {
	const version = "V1.3.1.241008_R"
	context.JSON(200, pkg.NewResponse(0, "Hello WPR!", map[string]any{
		"version": version,
	}))
}
