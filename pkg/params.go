package pkg

// SearchUsersParams
//
//	@Description: 查询用户参数
type SearchUsersParams struct {
	Page  *int // 页码
	Limit *int // 每页条数

	Account   *string // 账号
	Name      *string // 姓名
	Direction *int    // 方向
}

// SearchClockInParams
//
//	@Description: 查询打卡参数
type SearchClockInParams struct {
	Page  *int // 页码
	Limit *int // 每页条数

	UID      *int64    // 用户ID
	Date     *string   // 日期
	TimeSlot *TimeSlot // 时间段
}

type TimeSlot string

const (
	TimeSlotMorning   TimeSlot = "上午"
	TimeSlotAfternoon TimeSlot = "下午"
	TimeSlotNight     TimeSlot = "晚上"
	TimeSlotAllDay    TimeSlot = "全天"
)
