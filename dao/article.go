package dao

import (
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
)

type ArticleDao struct {
}

// GetArticleDetails
// @Date 2023-02-23 21:16:40
// @Param articleId int
// @Return article.ArticleDetails
// @Return error
// @Description: 获取一篇文章的详细信息
func (ArticleDao) GetArticleDetails(articleId int) (response.ArticleDetails, error) {
	var articleDetails response.ArticleDetails
	err := global.DB().
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
func (ArticleDao) GetArticleInfoList(pageInfo request.PageInfo) ([]response.HeadInfo, error) {
	page := PageHelper.SelectPage(pageInfo)
	var articleList []response.HeadInfo
	model := global.DB().Table("articles a").
		Select("a.id id", "u.nickname author", "a.title", "a.cover", "a.label", "a.summary", "a.view", "a.content", "a.updated_at").
		Joins("JOIN system_users u ON u.id = a.user_id")
	if err := page(model, nil, &articleList).Error; err != nil {
		return nil, err
	} else {
		return articleList, nil
	}
}

func (ArticleDao) GetArticleCommentList(pageInfo request.PageInfo) ([]response.ArticleCommentInfo, error) {
	//page := PageHelper.SelectPage(pageInfo)
	return nil, nil
}
