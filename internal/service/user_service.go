package service

import (
	"WorkProgressRecord/internal/service/impl"
	"WorkProgressRecord/pkg"
)

// UserService
//
//	@Description: 用户服务接口
type UserService interface {
	Login(account, password string) *pkg.Response
}

// NewUserService
//
//	@Description: 创建用户服务实例
//	@return UserService 用户服务实例
func NewUserService() UserService {
	return &impl.UserServiceImpl{}
}
