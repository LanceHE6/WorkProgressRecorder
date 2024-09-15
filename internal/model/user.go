package model

// User
//
//	@Description: 用户表结构
type User struct {
	Base
	Account    string `gorm:"column:account;type:varchar(255);not null;unique" json:"account"` // 账号
	Password   string `gorm:"column:password;type:varchar(255);not null" json:"-"`             // 密码
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`              // 姓名
	Permission int    `gorm:"column:permission;type:int;not null" json:"permission"`           // 权限
	SessionID  string `gorm:"column:session_id;type:varchar(255)" json:"-"`                    // session_id
}
