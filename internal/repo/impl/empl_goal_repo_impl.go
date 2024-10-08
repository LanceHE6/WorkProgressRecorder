package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

// EmplGoalRepoImpl
//
//	@Description: 就业目标仓库实现
type EmplGoalRepoImpl struct {
}

// modelDB
//
//	@Description: 获取模型数据库
//	@receiver *EmplGoalRepoImpl 就业目标仓库实现
//	@return *gorm.DB 模型数据库
func (*EmplGoalRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.EmploymentGoal{})
}

// Insert
//
//	@Description: 插入就业目标
//	@receiver e 就业目标仓库实现
//	@param emplGoal 就业目标
//	@return error 错误
func (e *EmplGoalRepoImpl) Insert(emplGoal model.EmploymentGoal) error {
	tx := db.GetMyDbConnection().Begin()
	err := tx.Model(&model.EmploymentGoal{}).Create(&emplGoal).Error
	if err != nil {
		return err
	}
	// 更新用户方向
	var user model.User
	err = tx.Model(&model.User{}).Where("id = ?", emplGoal.UID).First(&user).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	if !pkg.IsContainGoal(user.Direction, model.Employment) {
		user.Direction = append(user.Direction, model.Employment)
	}
	err = tx.Model(&model.User{}).Where("id = ?", emplGoal.UID).Update("direction", user.Direction).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// Update
//
//	@Description: 更新就业目标
//	@receiver e 就业目标仓库实现
//	@param emplGoal 就业目标
//	@return error 错误
func (e *EmplGoalRepoImpl) Update(emplGoal model.EmploymentGoal) error {
	tx := db.GetMyDbConnection().Begin()
	err := tx.Model(&model.EmploymentGoal{}).Update(&emplGoal).Error
	if err != nil {
		return err
	}
	// 更新用户方向
	var user model.User
	err = tx.Model(&model.User{}).Where("id = ?", emplGoal.UID).First(&user).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	if !pkg.IsContainGoal(user.Direction, model.Employment) {
		user.Direction = append(user.Direction, model.Employment)
	}
	err = tx.Model(&model.User{}).Where("id = ?", emplGoal.UID).Update("direction", user.Direction).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// SelectByUID
//
//	@Description: 根据用户id查询就业目标
//	@receiver e 就业目标仓库实现
//	@param uid 用户id
//	@return model.EmploymentGoal 就业目标
//	@return error 错误
func (e *EmplGoalRepoImpl) SelectByUID(uid int64) (model.EmploymentGoal, error) {
	var emplGoal model.EmploymentGoal
	err := e.modelDB().Where("uid = ?", uid).First(&emplGoal).Error
	return emplGoal, err
}
