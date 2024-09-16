package impl

import (
	"WorkProgressRecord/internal/model"
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
	err = tx.Model(&model.User{}).Where("id = ?", emplGoal.UID).Update("direction", 2).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
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
	return e.modelDB().Update(&emplGoal).Error
}
