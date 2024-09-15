package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

// UserRepository
// @Description: 用户仓库接口
type UserRepository interface {
	SelectByID(id int64) *model.User
	SelectByAccountAndPsw(account, password string) *model.User
}

// NewUserRepository
//
//	@Description: 创建用户仓库实例
//	@return UserRepository 用户仓库实例
func NewUserRepository() UserRepository {
	return &impl.UserRepositoryImpl{}
}
