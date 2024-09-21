package repo

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo/impl"
	"WorkProgressRecord/pkg"
)

type PostRepo interface {
	// Insert 创建一条新帖子
	Insert(post model.Post) error
	// AddComment 给帖子添加一条评论
	AddComment(pid int64, comment model.Comment) error
	// SelectByID 根据ID查询帖子
	SelectByID(id int64) (*model.Post, error)
	// Search 根据参数查询帖子
	Search(params pkg.SearchPostsParams) ([]model.Post, int, error)
	// DeleteByID 根据ID删除帖子
	DeleteByID(id int64) error
}

func NewPostRepo() PostRepo {
	return &impl.PostRepoImpl{}
}
