package model

type EmploymentGoal struct {
	Base
	// 关联用户
	User *User `gorm:"foreignKey:UID"`
	UID  int64 `gorm:"unique" json:"uid"`

	// 目标
	Status        int    `json:"status"`         // 就业状态1:未拿到offer 2:已拿到offer
	TargetCompany string `json:"target_company"` // 目标公司
	TargetJob     string `json:"target_job"`     // 目标职位
	TargetSalary  int    `json:"target_salary"`  // 目标薪资
	TargetArea    string `json:"target_area"`    // 目标地区
}
