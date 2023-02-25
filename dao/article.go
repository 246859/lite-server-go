package dao

import (
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
	"gorm.io/gorm"
)

type ArticleDao struct {
}

// GetArticleDetails
// @Date 2023-02-23 21:16:40
// @Param articleId int
// @Return article.ArticleDetails
// @Return error
// @Description: 获取一篇文章的详细信息
func (ArticleDao) GetArticleDetails(db *gorm.DB, articleId int) (response.ArticleDetails, error) {
	var articleDetails response.ArticleDetails
	err := db.
		Table("system_users u").
		Select("a.id id", "u.id user_id", "u.nickname author", "a.title", "a.cover", "a.label", "a.summary", "a.view", "a.content", "a.updated_at").
		Joins("JOIN articles a ON u.id = a.user_id").
		Where("a.id=?", articleId).Scan(&articleDetails).Error
	return articleDetails, err
}

// GetArticleInfoList
// @Date 2023-02-23 21:25:34
// @Param pageInfo sysreq.PageInfo
// @Return []article.HeadInfo
// @Return error
// @Description: 查询文章简单信息列表
func (ArticleDao) GetArticleInfoList(db *gorm.DB, pageInfo request.PageInfo) ([]response.HeadInfo, error) {
	page := PageHelper.SelectPage(pageInfo)
	var articleList []response.HeadInfo
	model := db.Table("articles a").
		Select("a.id id", "u.nickname author", "a.title", "a.cover", "a.label", "a.summary", "a.view", "a.content", "a.updated_at").
		Joins("JOIN system_users u ON u.id = a.user_id")
	if err := page(model, nil, &articleList).Error; err != nil {
		return nil, err
	} else {
		return articleList, nil
	}
}

// GetArticleCommentList
// @Date 2023-02-25 15:52:26
// @Param pageInfo request.PageInfo
// @Param articleId int
// @Return []response.ArticleCommentInfo
// @Return error
// @Description: 查询一个文章的评论信息
func (ArticleDao) GetArticleCommentList(db *gorm.DB, pageInfo request.PageInfo, articleId int) ([]response.ArticleCommentInfo, error) {
	var data []response.ArticleCommentInfo
	page := PageHelper.SelectPage(pageInfo)
	model := db.Table("comments c").
		Preload("User").
		Preload("ReplyList").
		Preload("ReplyList.User").
		Select("c.id", "c.user_id", "ac.article_id", "c.content", "c.updated_at").
		Joins("JOIN articles_comments ac ON ac.comment_id = c.id").
		Where("ac.article_id = ?", articleId)
	page(model, nil, &data)
	return data, model.Error
}
