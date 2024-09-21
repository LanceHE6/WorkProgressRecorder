package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Post
// @Description: 帖子
type Post struct {
	Base
	// 关联用户
	User *User `gorm:"foreignKey:UID" json:"-"`
	UID  int64 `gorm:"unique" json:"uid"`

	Tittle   string      `gorm:"type:varchar(255);not null" json:"tittle"`
	Content  string      `gorm:"type:text;not null" json:"content"`
	Comments CommentList `gorm:"type:json" json:"comments"`
}

type CommentList []Comment

type Comment struct {
	UID       int64  `json:"uid"`        // 状态
	CreatedAt int64  `json:"created_at"` // 创建时间
	Content   string `json:"content"`    // 内容
}

//  实现Scanner和Valuer接口

func (s *CommentList) Scan(value interface{}) error {
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

func (s CommentList) Value() (driver.Value, error) {
	// 将StatusTimeLine编码为JSON格式的[]byte
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
