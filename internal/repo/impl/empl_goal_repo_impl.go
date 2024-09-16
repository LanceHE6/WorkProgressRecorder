package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

type EmplGoalRepoImpl struct {
}

func (*EmplGoalRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.EmploymentGoal{})
}

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
