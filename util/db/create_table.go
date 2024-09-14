package db

import (
	"WorkProgressRecord/entity/model"
	"github.com/jinzhu/gorm"
)

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
