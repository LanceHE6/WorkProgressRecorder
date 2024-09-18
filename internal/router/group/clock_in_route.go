package group

import (
	"WorkProgressRecord/internal/handler"
	"WorkProgressRecord/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ClockInRoute(group *gin.RouterGroup) {
	clockInHandler := handler.ClockInHandler{}

	clockInGroup := group.Group("/clock_in")
	clockInGroup.Use(middleware.AuthMiddleware())
	clockInGroup.POST("/add", func(context *gin.Context) {
		clockInHandler.AddClockIn(context)
	})
}
