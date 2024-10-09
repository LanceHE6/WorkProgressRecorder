package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Post
// @Description: 帖子
type Post struct {
	Base
	// 关联用户
	User *User `gorm:"foreignKey:UID" json:"user"`
	UID  int64 `json:"uid"`

	Title     string      `gorm:"type:text;not null" json:"title"`
	Content   string      `gorm:"type:text;not null" json:"content"`
	Anonymous bool        `gorm:"type:tinyint(1);not null" json:"anonymous"`
	Comments  CommentList `gorm:"type:json" json:"comments"`
}

type CommentList []Comment

type Comment struct {
	User      User      `json:"user"`       // 评论用户
	CreatedAt time.Time `json:"created_at"` // 创建时间
	Content   string    `json:"content"`    // 内容
	Anonymous bool      `json:"anonymous"`  // 是否匿名
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
