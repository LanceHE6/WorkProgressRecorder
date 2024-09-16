package model

type PostgraduateGoal struct {
	Base
	// 关联用户
	User User  `gorm:"foreignKey:UID" json:"-"`
	UID  int64 `gorm:"unique" json:"uid"`

	TargetUniversity string  `json:"target_university"` // 目标院校
	TargetMajor      string  `json:"target_major"`      // 目标专业
	TargetScore      float64 `json:"target_score"`      // 目标分数
}
