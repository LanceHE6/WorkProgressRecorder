package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

// EmplGoalService
//
//	@Description: 就业目标服务接口
type EmplGoalService interface {
	InsertAndUpdate(context *gin.Context)
}

// NewEmplGoalService
//
//	@Description: 创建就业目标服务实例
//	@return EmplGoalService 就业目标服务实例
func NewEmplGoalService() EmplGoalService {
	return &impl.EmplGoalServiceImpl{}
}
