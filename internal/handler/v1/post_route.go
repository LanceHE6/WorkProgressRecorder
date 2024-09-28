package v1

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
)

// PostRoute
//
//	@Description: 帖子路由
//	@param group *gin.RouterGroup
func PostRoute(group *gin.RouterGroup) {
	postService := service.NewPostService()

	postGroup := group.Group("/post")
	postGroup.GET("/search", func(context *gin.Context) {
		postService.SearchPosts(context)
	})

	postGroup.Use(middleware.AuthMiddleware())
	postGroup.POST("/add", func(context *gin.Context) {
		postService.AddPost(context)
	})
	postGroup.POST("/comment/add", func(context *gin.Context) {
		postService.AddComment(context)
	})
	postGroup.DELETE("/del", func(context *gin.Context) {
		postService.DeletePost(context)
	})

}
