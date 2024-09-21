package model

type ClockIn struct {
	Base
	// 关联用户
	User            *User  `gorm:"foreignKey:UID" json:"user"`
	UID             int64  `json:"uid"`
	ClockInTime     int64  `json:"clock_in_time"`     // 打卡时间戳
	ClockInLocation string `json:"clock_in_location"` // 打卡地点
}
