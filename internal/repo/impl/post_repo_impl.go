package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/pkg"
	"WorkProgressRecord/pkg/db"
	"github.com/jinzhu/gorm"
)

type PostRepoImpl struct {
}

func (PostRepoImpl) modelDB() *gorm.DB {
	return db.GetMyDbConnection().Model(&model.Post{})
}

// Insert
//
//	@Description: 插入帖子
//	@receiver p PostRepoImpl
//	@param post 帖子
//	@return error 错误
func (p PostRepoImpl) Insert(post model.Post) error {
	return p.modelDB().Create(&post).Error
}

// AddComment
//
//	@Description: 添加评论
//	@receiver p PostRepoImpl
//	@param pid 帖子id
//	@param comment 评论
//	@return error 错误
func (p PostRepoImpl) AddComment(pid int64, comment model.Comment) error {
	post, err := p.SelectByID(pid)
	if err != nil {
		return err
	}
	post.Comments = append(post.Comments, comment)
	return p.modelDB().Update(post).Error
}

// SelectByID
//
//	@Description: 根据id查询帖子
//	@receiver p PostRepoImpl
//	@param id 帖子id
//	@return *model.Post 帖子
//	@return error 错误
func (p PostRepoImpl) SelectByID(id int64) (*model.Post, error) {
	var post model.Post
	err := p.modelDB().Where("id = ?", id).First(&post).Error
	return &post, err
}

// Search
//
//	@Description: 搜索帖子
//	@receiver p PostRepoImpl
//	@param params 搜索参数
//	@return []model.Post 帖子列表
//	@return int 总数
//	@return error 错误
func (p PostRepoImpl) Search(params pkg.SearchPostsParams) ([]model.Post, int, error) {
	var posts []model.Post
	query := p.modelDB().Preload("User")
	// 构建查询
	if params.UID != nil {
		query = query.Where("uid = ?", *params.UID)
	}
	// 计算总数
	var count int
	query.Count(&count)
	// 分页
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	} else {
		params.Limit = new(int)
		*params.Limit = 10
	}
	if params.Page != nil {
		query = query.Offset((*params.Page - 1) * *params.Limit)
	}
	// 执行查询
	err := query.Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	return posts, count, nil
}

// DeleteByID
//
//	@Description: 根据id删除帖子
//	@receiver p PostRepoImpl
//	@param id 帖子id
//	@return error 错误
func (p PostRepoImpl) DeleteByID(id int64) error {
	return p.modelDB().Where("id = ?", id).Delete(&model.Post{}).Error
}
