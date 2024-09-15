package model

import "time"

// Base 基础模型，所有模型都继承该模型
type Base struct {
	ID        int64     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
