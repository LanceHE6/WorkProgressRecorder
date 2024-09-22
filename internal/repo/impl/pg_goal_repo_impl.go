package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

// PGGoalRepoImpl
//
//	@Description: 考研目标仓库实现
type PGGoalRepoImpl struct {
}

// modelDB
//
//	@Description: 获取模型数据库
//	@receiver *PGGoalRepoImpl 考研目标仓库实现
//	@return *gorm.DB 模型数据库
func (*PGGoalRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.PostgraduateGoal{})
}

// Insert
//
//	@Description: 插入考研目标
//	@receiver PGGoalRepoImpl 考研目标仓库实现
//	@param pgGoal 考研目标
//	@return error 错误
func (p *PGGoalRepoImpl) Insert(pgGoal model.PostgraduateGoal) error {
	tx := db.GetMyDbConnection().Begin()
	err := tx.Model(&model.PostgraduateGoal{}).Create(&pgGoal).Error
	if err != nil {
		return err
	}
	// 更新用户方向
	var user model.User
	err = tx.Model(&model.User{}).Where("id = ?", pgGoal.UID).First(&user).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	if !pkg.IsContainGoal(user.Direction, model.Postgraduate) {
		user.Direction = append(user.Direction, model.Postgraduate)
	}
	err = tx.Model(&model.User{}).Where("id = ?", pgGoal.UID).Update("direction", user.Direction).Error
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
//	@Description: 更新考研目标
//	@receiver PGGoalRepoImpl 考研目标仓库实现
//	@param pgGoal 考研目标
//	@return error 错误
func (p *PGGoalRepoImpl) Update(pgGoal model.PostgraduateGoal) error {
	tx := db.GetMyDbConnection().Begin()
	err := tx.Model(&model.PostgraduateGoal{}).Updates(&pgGoal).Error
	if err != nil {
		return err
	}
	// 更新用户方向
	var user model.User
	err = tx.Model(&model.User{}).Where("id = ?", pgGoal.UID).First(&user).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	if !pkg.IsContainGoal(user.Direction, model.Postgraduate) {
		user.Direction = append(user.Direction, model.Postgraduate)
	}
	err = tx.Model(&model.User{}).Where("id = ?", pgGoal.UID).Update("direction", user.Direction).Error
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
//	@Description: 根据用户id查询考研目标
//	@receiver p *PGGoalRepoImpl
//	@param uid 用户id
//	@return model.PostgraduateGoal 考研目标
//	@return error 错误
func (p *PGGoalRepoImpl) SelectByUID(uid int64) (model.PostgraduateGoal, error) {
	var pgGoal model.PostgraduateGoal
	err := p.modelDB().Where("uid = ?", uid).First(&pgGoal).Error
	return pgGoal, err
}
