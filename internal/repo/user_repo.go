package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
	"WorkProgressRecord/pkg"
)

// UserRepository
// @Description: 用户仓库接口
type UserRepository interface {
	SelectByID(id int64) *model.User
	SelectByAccountAndPsw(account, password string) *model.User
	Insert(user model.User) error
	UpdateSessionID(id int64, sessionID string) error
	UpdateDirection(id int64, direction int) error
	UpdatePassword(id int64, newPsw string) error
	SearchUsers(params pkg.SearchUsersParams) ([]model.User, int)
}

// NewUserRepository
//
//	@Description: 创建用户仓库实例
//	@return UserRepository 用户仓库实例
func NewUserRepository() UserRepository {
	return &impl.UserRepositoryImpl{}
}
