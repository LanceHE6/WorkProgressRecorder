package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// User
//
//	@Description: 用户表结构
type User struct {
	Base
	Account    string    `gorm:"column:account;type:varchar(255);not null;unique" json:"account"` // 账号
	Password   string    `gorm:"column:password;type:varchar(255);not null" json:"-"`             // 密码
	Name       string    `gorm:"column:name;type:varchar(255);not null" json:"name"`              // 姓名
	Class      string    `gorm:"column:class;type:varchar(255);" json:"class"`                    // 班级
	Major      string    `gorm:"column:major;type:varchar(255);" json:"major"`                    // 专业
	Permission int       `gorm:"column:permission;type:int;not null" json:"permission"`           // 权限
	Direction  Direction `gorm:"column:direction;type:json;" json:"direction"`                    // 数组json对象,方向1:考研 2:就业
	SessionID  string    `gorm:"column:session_id;type:varchar(255)" json:"-"`                    // session_id
}

// 方向定义

type Direction []DirectionType

type DirectionType int

const (
	Postgraduate DirectionType = iota + 1
	Employment   DirectionType = iota + 1
)

//  实现Scanner和Valuer接口

func (s *Direction) Scan(value interface{}) error {
	// 假设value是一个[]byte（从数据库读取的JSON数据）
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to cast value to []byte")
	}

	// 将JSON数据解码到s指向的StatusTimeLine实例中
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}

	return nil
}

func (s Direction) Value() (driver.Value, error) {
	// 将StatusTimeLine编码为JSON格式的[]byte
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
