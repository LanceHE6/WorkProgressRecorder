package service

import (
	"WorkProgressRecord/internal/service/impl"
	"github.com/gin-gonic/gin"
)

type PostService interface {
	AddPost(context *gin.Context)
	AddComment(context *gin.Context)
	SearchPosts(context *gin.Context)
	DeletePost(context *gin.Context)
}

func NewPostService() PostService {
	return &impl.PostServiceImpl{}
}
