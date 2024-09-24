package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
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

// Insert
//
//	@Description: 插入用户
//	@receiver u UserRepositoryImpl
//	@param user *model.User 用户数据
//	@return error 错误信息
func (u UserRepositoryImpl) Insert(user model.User) error {
	// 密码加密
	// 密码默认未账号后6位
	user.Password = pkg.HashPsw(user.Account[len(user.Account)-6:])
	user.Permission = 1
	return u.modelDB().Create(&user).Error
}

// UpdateSessionID
//
//	@Description: 更新用户sessionID
//	@receiver u UserRepositoryImpl
//	@param id int64 用户id
//	@param sessionID string sessionID
//	@return error 错误信息
func (u UserRepositoryImpl) UpdateSessionID(id int64, sessionID string) error {
	return u.modelDB().Where("id = ?", id).Update("session_id", sessionID).Error
}

// UpdateDirection
//
//	@Description: 更新用户方向
//	@receiver u UserRepositoryImpl
//	@param id int64 用户id
//	@param direction int 方向
//	@return error 错误信息
func (u UserRepositoryImpl) UpdateDirection(id int64, direction model.Direction) error {
	return u.modelDB().Where("id = ?", id).Update("direction", direction).Error
}

// UpdatePassword
//
//	@Description: 更新用户密码
//	@receiver u UserRepositoryImpl
//	@param id int64 用户id
//	@param newPsw string 新密码
//	@return error 错误信息
func (u UserRepositoryImpl) UpdatePassword(id int64, newPsw string) error {
	return u.modelDB().Where("id = ?", id).Update("password", pkg.HashPsw(newPsw)).Error
}

func (u UserRepositoryImpl) SearchUsers(params pkg.SearchUsersParams) ([]model.User, int) {
	query := db.GetMyDbConnection().Table("users")
	if params.Account != nil {
		query = query.Where("account like ?", "%"+*params.Account+"%")
	}
	if params.Name != nil {
		query = query.Where("name like ?", "%"+*params.Name+"%")
	}
	if params.Keyword != nil {
		query = query.Where("account like ? OR name like ? ",
			"%"+*params.Keyword+"%",
			"%"+*params.Keyword+"%",
		)
	}
	if params.Direction != nil {
		// 未填写方向的
		if *params.Direction == 0 {
			query = query.Where("JSON_LENGTH(direction) = 0")
		} else // 填写了两个方向的
		if *params.Direction == 3 {
			query = query.Where("JSON_CONTAINS(direction, '1') AND JSON_CONTAINS(direction, '2')")
		} else {
			// 查询单个方向的
			query = query.Where("JSON_CONTAINS(direction, ?)", strconv.Itoa(*params.Direction))
		}
	}
	// 统计总数
	var count int
	query.Count(&count)
	// 分页
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	} else {
		params.Limit = new(int)
		*params.Limit = 10
	}
	if params.Page != nil {
		query = query.Offset((*params.Page - 1) * *params.Limit)
	}
	var users []model.User
	query.Find(&users)
	return users, count
}

// UpdateUserInfo
//
//	@Description: 更新用户信息
//	@receiver u UserRepositoryImpl
//	@param id int64
//	@param user model.User
//	@return error 错误信息
func (u UserRepositoryImpl) UpdateUserInfo(id int64, user model.User) error {
	return u.modelDB().Where("id = ?", id).Updates(user).Error
}

// DeleteDirection
//
//	@Description: 删除用户方向
//	@receiver u UserRepositoryImpl
//	@param id int64 用户id
//	@param direction model.DirectionType 方向
//	@return error 错误信息
func (u UserRepositoryImpl) DeleteDirection(id int64, direction model.DirectionType) error {
	user := u.SelectByID(id)
	if pkg.IsContainGoal(user.Direction, direction) {
		user.Direction = pkg.RemoveGoal(user.Direction, direction)
		return u.UpdateDirection(id, user.Direction)
	} else {
		return errors.New("不能删除用户没有的方向")
	}
}
