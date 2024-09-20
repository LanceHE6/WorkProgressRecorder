package db

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// InitData
//
//	@Description: 初始化数据
func InitData() {
	// 初始化用户数据
	createIfNotExists(db, &model.User{
		Account:    "admin",
		Password:   pkg.HashPsw("admin"),
		Name:       "admin",
		Permission: 2,
	}, "admin", "account")
	// 初始化就业状态数据
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "期望投递",
	}, "期望投递", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "已听宣讲",
	}, "已听宣讲", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "已经投递",
	}, "已经投递", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "等待面试",
	}, "等待面试", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "已经面试",
	}, "已经面试", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "正在面试(多轮面试)",
	}, "正在面试(多轮面试)", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "失败",
	}, "失败", "status_name")
	createIfNotExists(db, &model.WorkProgressStatus{
		StatusName: "拿到offer",
	}, "拿到offer", "status_name")
}

// createIfNotExists
//
//	@Description: 创建数据，如果不存在
//	@param db 数据库
//	@param value 数据
//	@param id id
//	@param idFieldName id字段名
func createIfNotExists(db *gorm.DB, value interface{}, id interface{}, idFieldName string) {
	// 检查数据是否存在
	if err := db.Where(idFieldName+" = ?", id).First(value).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 数据不存在，插入新数据
			if err = db.Create(value).Error; err != nil {
				// 插入错误
				log.Println(err)
				os.Exit(-100)
			}
		} else {
			// 查询错误
			log.Println(err)
			os.Exit(-200)
		}
	}
}
