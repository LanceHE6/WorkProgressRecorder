package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
)

// WorkLogRepo
//
//	@Description: 工作日志仓库接口
type WorkLogRepo interface {
	InsertWorkLog(workLog model.WorkLog) error
	SelectByID(id int64) (model.WorkLog, error)
	AddStatusTimeLine(id int64, timeLine model.StatusTime) error
	SelectByUID(uid int64) ([]model.WorkLog, error)
}

// NewWorkLogRepo
//
//	@Description: 创建工作日志仓库实例
//	@return WorkLogRepo 工作日志仓库实例
func NewWorkLogRepo() WorkLogRepo {
	return &impl.WorkLogRepoImpl{}
}
