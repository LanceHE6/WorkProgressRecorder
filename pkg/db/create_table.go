package db

import (
	"WorkProgressRecord/internal/model"
	"github.com/jinzhu/gorm"
)

// CreateTable
//
//	@Description: 创建表
//	@param db *gorm.DB 数据库连接
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
