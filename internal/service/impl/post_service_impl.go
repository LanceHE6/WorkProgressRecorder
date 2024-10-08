package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type PostServiceImpl struct {
}

// AddPost
//
//	@Description: 添加帖子
//	@receiver p PostServiceImpl
//	@param context *gin.Context
func (p PostServiceImpl) AddPost(context *gin.Context) {
	type addPostReq struct {
		Title     string `json:"title" form:"title" binding:"required"`
		Content   string `json:"content" form:"content" binding:"required"`
		Anonymous bool   `json:"anonymous" form:"anonymous"`
	}
	var data addPostReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}

	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}

	post := model.Post{
		UID:       userInfo.ID,
		Title:     data.Title,
		Content:   data.Content,
		Anonymous: data.Anonymous,
	}

	postRepo := repo.NewPostRepo()
	if err = postRepo.Insert(post); err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(1, "添加失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// AddComment
//
//	@Description: 添加评论
//	@receiver p PostServiceImpl
//	@param context *gin.Context
func (p PostServiceImpl) AddComment(context *gin.Context) {
	type addCommentReq struct {
		ID        int64  `json:"id" form:"id" binding:"required"`
		Content   string `json:"content" form:"content" binding:"required"`
		Anonymous bool   `json:"anonymous" form:"anonymous"`
	}
	var data addCommentReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	userRepo := repo.NewUserRepository()
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}

	// 获取评论用户实列
	user := userRepo.SelectByID(userInfo.ID)
	if user == nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-2, "用户不存在", err))
		return
	}
	comment := model.Comment{
		User:      *user,
		Content:   data.Content,
		CreatedAt: time.Now(),
		Anonymous: data.Anonymous,
	}

	postRepo := repo.NewPostRepo()
	if err = postRepo.AddComment(data.ID, comment); err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(1, "添加失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// SearchPosts
//
//	@Description: 搜索帖子
//	@receiver p PostServiceImpl
//	@param context *gin.Context
func (p PostServiceImpl) SearchPosts(context *gin.Context) {
	var params pkg.SearchPostsParams
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	params.Page = &page
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	params.Limit = &limit
	if context.Query("uid") != "" {
		uid, _ := strconv.Atoi(context.Query("uid"))
		uid64 := int64(uid)
		params.UID = &uid64
	}
	postRepo := repo.NewPostRepo()
	postsList, count, err := postRepo.Search(params)
	if err != nil {
		context.JSON(http.StatusOK, pkg.ErrorResponse(-1, "查询失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"rows":      postsList,
		"total":     count,
		"page":      page,
		"page_size": limit,
	}))
}

// DeletePost
//
//	@Description: 删除帖子
//	@receiver p PostServiceImpl
//	@param context *gin.Context
func (p PostServiceImpl) DeletePost(context *gin.Context) {
	type deletePostReq struct {
		ID int64 `json:"id" form:"id" binding:"required"`
	}
	var data deletePostReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	postRepo := repo.NewPostRepo()
	post, err := postRepo.SelectByID(data.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "不存在的帖子", err))
		return
	}
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-2, "无法获取用户id", err))
	}
	if post.UID != userInfo.ID {
		context.JSON(http.StatusOK, pkg.FailedResponse(1, "无权限删除"))
		return
	}
	if err = postRepo.DeleteByID(data.ID); err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-3, "删除失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}
