package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

type WPStatusRepoImpl struct {
}

func (*WPStatusRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.WorkProgressStatus{})
}

// ListStatus
//
//	@Description: 获取所有状态
//	@receiver w *WPStatusRepoImpl
//	@return []model.WorkProgressStatus 状态列表
//	@return error 错误信息
func (w *WPStatusRepoImpl) ListStatus() ([]model.WorkProgressStatus, error) {
	var status []model.WorkProgressStatus
	err := w.modelDB().Find(&status).Error
	return status, err
}
