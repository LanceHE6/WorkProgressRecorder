package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

// UserService
//
//	@Description: 用户服务接口
type UserService interface {
	Login(context *gin.Context)
	Import(context *gin.Context)
	UpdatePassword(context *gin.Context)
	GetUserTargetInfo(context *gin.Context)
	GetUserInfo(context *gin.Context)
	SearchUsers(context *gin.Context)
	UpdateUserInfo(context *gin.Context)
}

// NewUserService
//
//	@Description: 创建用户服务实例
//	@return UserService 用户服务实例
func NewUserService() UserService {
	return &impl.UserServiceImpl{}
}
