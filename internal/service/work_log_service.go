package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

type WorkLogService interface {
	AddWorkLog(context *gin.Context)
	AddStatusTimeLine(context *gin.Context)
	GetUserWorkLog(context *gin.Context)
}

func NewWorkLogService() WorkLogService {
	return &impl.WorkLogServiceImpl{}
}
