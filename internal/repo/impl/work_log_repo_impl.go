package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

type WorkLogRepoImpl struct {
}

func (WorkLogRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.WorkLog{})
}

// InsertWorkLog
//
//	@Description: 插入一条工作日志
//	@receiver w WorkLogRepoImpl
//	@param workLog 工作日志
//	@return error 错误信息
func (w WorkLogRepoImpl) InsertWorkLog(workLog model.WorkLog) error {
	return w.modelDB().Create(&workLog).Error
}

// SelectByID
//
//	@Description: 根据id查询一条工作日志
//	@receiver w WorkLogRepoImpl
//	@param id
//	@return model.WorkLog 工作日志
//	@return error 错误信息
func (w WorkLogRepoImpl) SelectByID(id int64) (model.WorkLog, error) {
	var workLog model.WorkLog
	err := w.modelDB().Where("id=?", id).First(&workLog).Error
	return workLog, err
}

// AddStatusTimeLine
//
//	@Description: 添加一条状态时间线
//	@receiver w WorkLogRepoImpl
//	@param id 工作日志id
//	@param timeLine 状态时间线
//	@return error 错误信息
func (w WorkLogRepoImpl) AddStatusTimeLine(id int64, timeLine model.StatusTime) error {
	workLog, err := w.SelectByID(id)
	if err != nil {
		return err
	}
	newLine := append(workLog.StatusTimeLine, timeLine)
	workLog.StatusTimeLine = newLine
	return w.modelDB().Update(&workLog).Error
}

// SelectByUID
//
//	@Description: 根据用户id查询工作日志
//	@receiver w WorkLogRepoImpl
//	@param uid 用户id
//	@return []model.WorkLog 工作日志列表
//	@return error 错误信息
func (w WorkLogRepoImpl) SelectByUID(uid int64) ([]model.WorkLog, error) {
	var workLogs []model.WorkLog
	err := w.modelDB().Where("uid=?", uid).Find(&workLogs).Error
	return workLogs, err
}
