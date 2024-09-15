package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

// UserRepositoryImpl
//
//	@Description: 用户仓库实现
type UserRepositoryImpl struct {
}

// modelDB
//
//	@Description: 获取用户表
//	@receiver u UserRepositoryImpl
//	@return *gorm.DB 用户表
func (u UserRepositoryImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(model.User{})
}

// SelectByID
//
//	@Description: 根据id查询用户
//	@receiver u UserRepositoryImpl
//	@param id int64 用户id
//	@return *model.User 用户数据
func (u UserRepositoryImpl) SelectByID(id int64) *model.User {
	var user model.User
	err := u.modelDB().Where("id = ?", id).First(&user)
	if err.Error != nil {
		return nil
	}
	return &user
}

// SelectByAccountAndPsw
//
//	@Description: 根据账号密码查询用户
//	@receiver u UserRepositoryImpl
//	@param account string 账号
//	@param password string 密码
//	@return *model.User 用户数据
func (u UserRepositoryImpl) SelectByAccountAndPsw(account, password string) *model.User {
	var user model.User
	err := u.modelDB().Where("account = ? and password = ?", account, pkg.HashPsw(password)).First(&user)
	if err.Error != nil {
		return nil
	}
	return &user
}
