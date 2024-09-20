package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

// ClockInRepoImpl
//
//	@Description: 打卡记录仓库实现
type ClockInRepoImpl struct {
}

// modelDB
//
//	@Description: 获取打卡记录表
//	@receiver *ClockInRepoImpl 打卡记录仓库
//	@return *gorm.DB 打卡记录表
func (*ClockInRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.ClockIn{})
}

// Insert
//
//	@Description: 插入打卡记录
//	@receiver c 打卡记录仓库
//	@param clockInData 打卡记录
//	@return error 错误信息
func (c *ClockInRepoImpl) Insert(clockInData model.ClockIn) error {
	return c.modelDB().Create(&clockInData).Error
}

// SelectUserLatest
//
//	@Description: 查询用户最新的一条打卡记录
//	@receiver c 打卡记录仓库
//	@param uid 用户id
//	@return model.ClockIn 打卡记录
//	@return error 错误信息
func (c *ClockInRepoImpl) SelectUserLatest(uid int64) (model.ClockIn, error) {
	var clockIn model.ClockIn
	err := c.modelDB().Where("uid=?", uid).Order("clock_in_time desc").First(&clockIn).Error
	return clockIn, err
}
