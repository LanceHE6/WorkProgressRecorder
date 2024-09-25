package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

type VersionService interface {
	GetVersion(context *gin.Context)
}

func NewVersionService() VersionService {
	return &impl.VersionServiceImpl{}
}
