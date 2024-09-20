package main

import (
	"WorkProgressRecord/config"
	"WorkProgressRecord/internal/handler"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	// 跨域
	ginServer.Use(pkg.Cors())

	// 加载路由
	handler.Route(ginServer)

	err := ginServer.Run(":" + config.ServerConfig.SERVER.PORT)
	if err != nil {
		return
	}
}
