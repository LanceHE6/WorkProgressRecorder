package main

import (
	"WorkProgressRecord/util"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.Use(util.Cors())

	err := ginServer.Run("8070")
	if err != nil {
		return
	}
}
