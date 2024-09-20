package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

// PGGoalService
//
//	@Description: 考研目标服务接口
type PGGoalService interface {
	InsertAndUpdate(context *gin.Context)
}

// NewPGGoalService
//
//	@Description: 创建考研目标服务实例
//	@return PGGoalService 考研目标服务实例
func NewPGGoalService() PGGoalService {
	return &impl.PGGoalServiceImpl{}
}
