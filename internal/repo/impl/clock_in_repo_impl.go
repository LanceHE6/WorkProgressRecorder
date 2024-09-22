package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"fmt"
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

// SearchClockIns
//
//		@Description: 查询符合指定UID,日期和时间段的打卡记录
//		@receiver c 打卡记录仓库
//		@param params 查询参数
//		@return []model.ClockIn 打卡记录列表
//		@return error 错误信息
//	 @return int 总数
func (c *ClockInRepoImpl) SearchClockIns(params pkg.SearchClockInParams) ([]model.ClockIn, int, error) {
	var clockIns []model.ClockIn
	query := c.modelDB().Preload("User")
	// 构建查询
	if params.UID != nil {
		query = query.Where("uid = ?", *params.UID)
	}
	if params.Date != nil {
		query = query.Where("date(created_at) = ?", *params.Date)
	}
	if params.TimeSlot != nil {
		switch *params.TimeSlot {
		case pkg.TimeSlotMorning:
			query = query.Where("hour(created_at) >= 4 AND hour(created_at) < 12")
		case pkg.TimeSlotAfternoon:
			query = query.Where("hour(created_at) >= 13 AND hour(created_at) < 18")
		case pkg.TimeSlotNight:
			query = query.Where("(hour(created_at) >= 19 OR hour(created_at) < 3)")
		case pkg.TimeSlotAllDay:
			break
		default:
			// 处理错误或者默认情况
			return nil, 0, fmt.Errorf("invalid time slot")
		}
	}
	// 计算总数
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
	// 执行查询
	query = query.Order("created_at desc")
	err := query.Find(&clockIns).Error
	if err != nil {
		return nil, 0, err
	}

	return clockIns, count, nil
}
