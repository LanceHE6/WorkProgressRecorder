package db

import (
	"WorkProgressRecord/config"
	"WorkProgressRecord/pkg"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB

// Init
//
//	@Description: 初始化数据库连接
func init() {
	var err error
	account := config.GetDBMySQLAccount()
	password := config.GetDBMySQLPassword()
	host := config.GetDBMySQLHost()
	port := config.GetDBMySQLPort()
	dbname := config.GetDBMySQLDBName()

	// 创建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		account,
		password,
		host,
		port,
	)

	// 连接到MySQL
	tdb, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("can not connect to database: " + err.Error())
		os.Exit(-1)
		return
	}

	// 创建数据库
	_, err = tdb.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		fmt.Println("can not create database: " + err.Error())
		os.Exit(-2)
		return
	}
	// 关闭数据库连接
	err = tdb.Close()
	if err != nil {
		fmt.Println("can not close database: " + err.Error())
		os.Exit(-3)
		return
	}

	// 创建MySQL连接字符串

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		account,
		password,
		host,
		port,
		dbname,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Cannot connect to MYSQL database: %v", err)
		os.Exit(-4)
	}

	// 初始化雪花算法生成worker
	worker, err := pkg.NewWorker(1)
	if err != nil {
		log.Println("Snow Flake NewWorker error: ", err)
		return
	}
	// 设置雪花算法生成id
	db.Callback().Create().Before("gorm:create").Replace("id", func(scope *gorm.Scope) {
		id := worker.NextId()
		err := scope.SetColumn("id", id)
		if err != nil {
			log.Println("Cannot set snowflake id: ", err)
			return
		}
	})

	// 自动建表
	CreateTable(db)

	// 初始化数据
	fmt.Println("Init data...")
	InitData()
	fmt.Println("Init data done.")
}

// CloseMyDb
//
//	@Description: 关闭数据库连接
func CloseMyDb() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Println("Close Db error: ", err)
		}
	}
}

// GetMyDbConnection
//
//	@Description: 获取数据库连接
//	@return *gorm.DB 数据库连接
func GetMyDbConnection() *gorm.DB {
	return db
}
