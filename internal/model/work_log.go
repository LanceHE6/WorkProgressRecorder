package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// WorkLog
//
//	@Description: 工作日志表结构
type WorkLog struct {
	Base
	// 关联用户
	User           *User          `gorm:"foreignKey:UID" json:"-"`
	UID            int64          `json:"uid"`
	CompanyName    string         `json:"company_name" gorm:"column:company_name;type:varchar(255);not null;"` // 公司名称
	Job            string         `json:"job" gorm:"column:job;type:varchar(255);not null;"`                   // 职位
	Salary         string         `json:"salary" gorm:"column:salary;type:varchar(255);not null;"`             // 薪资
	Location       string         `json:"location" gorm:"column:location;type:varchar(255);not null;"`         // 工作地点
	StatusTimeLine StatusTimeLine `gorm:"type:json" json:"status_time_line"`
}

// StatusTime
//
//	@Description: 工作进度状态时间线结构
type StatusTime struct {
	CreatedAt int64  `json:"created_at"` // 创建时间
	Status    string `json:"status"`     // 状态
}

type StatusTimeLine []StatusTime

//  实现Scanner和Valuer接口

func (s *StatusTimeLine) Scan(value interface{}) error {
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

func (s StatusTimeLine) Value() (driver.Value, error) {
	// 将StatusTimeLine编码为JSON格式的[]byte
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
