package main

import (
	"WorkProgressRecord/util"
	"WorkProgressRecord/util/config"
	"WorkProgressRecord/util/db"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	// 跨域
	ginServer.Use(util.Cors())
	// 初始化数据库
	db.Init()
	err := ginServer.Run(":" + config.ServerConfig.SERVER.PORT)
	if err != nil {
		return
	}
}
