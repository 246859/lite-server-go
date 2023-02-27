package article

import (
	"errors"
	"github.com/246859/lite-server-go/dao"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/model/response"
)

type ArticleService struct {
	classDao dao.ClassDao
}

var ArticleDao = new(dao.ArticleDao)

// Article
// @Date 2023-02-20 16:25:22
// @Param articleId string
// @Return *article.Article
// @Return error
// @Description: 根据单个id查询文章内容
func (a ArticleService) Article(articleId int) (*response.ArticleDetails, error) {
	if articleId < 0 {
		return nil, errors.New(global.I18nRawCN("article.invalidId"))
	}
	articleDetails, err := ArticleDao.GetArticleDetails(global.DB(), articleId)
	if err != nil {
		return nil, err
	} else {
		return &articleDetails, nil
	}
}

// ArticlePage
// @Date 2023-02-20 17:36:17
// @Param pageInfo sysreq.PageInfo
// @Return []article.HeadInfo
// @Return error
// @Description: 分页查询文章信息列表
func (a ArticleService) ArticlePage(pageInfo request.PageInfo) ([]response.HeadInfo, error) {
	list, err := ArticleDao.GetArticleInfoList(global.DB(), pageInfo)
	return list, err
}

// ArticleCommentList
// @Date 2023-02-25 15:57:03
// @Param pageInfo request.PageInfo
// @Param articleId int
// @Return []response.ArticleCommentInfo
// @Return error
// @Description: 分页查询一个文章的所有信息列表
func (a ArticleService) ArticleCommentList(pageInfo request.PageInfo, articleId int) ([]response.ArticleCommentInfo, error) {
	list, err := ArticleDao.GetArticleCommentList(global.DB(), pageInfo, articleId)
	return list, err
}

// CreateArticle
// @Date 2023-02-20 16:46:19
// @Param article *article.Article
// @Return error
// @Description: 创建一篇新文章
func (a ArticleService) CreateArticle(newArticle *request.PostArticle, userId uint) error {
	var article model.Article
	// 赋值
	article.UserId = userId
	article.Title = newArticle.Title
	article.Cover = newArticle.Cover
	article.Label = newArticle.Label
	article.Summary = newArticle.Summary
	article.Content = newArticle.Content
	class, err := a.classDao.GetOne(global.DB(), newArticle.Class)
	if err != nil {
		return err
	}
	// 查询分类ID
	article.ClassId = class.ID
	return global.DB().Model(model.Article{}).Create(&article).Error
}

// DeleteArticle
// @Date 2023-02-20 16:51:16
// @Param id string
// @Return error
// @Description: 删除文章
func (a ArticleService) DeleteArticle(id int) error {
	err := global.DB().Model(model.Article{}).Delete("id=?", id).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateArticle
// @Date 2023-02-20 17:00:38
// @Param articleOld *article.Article
// @Return error
// @Description: 更新文章
func (a ArticleService) UpdateArticle(articleOld *model.Article) error {
	if !a.HasArticle(articleOld.ID) {
		return errors.New(global.I18nRawCN("article.notExist"))
	}
	err := global.Model(model.Article{}).Updates(articleOld).Error
	if err != nil {
		return err
	}
	return nil
}

// HasArticle
// @Date 2023-02-20 16:58:12
// @Param id uint
// @Return bool
// @Description: 判断文章是否存在
func (a ArticleService) HasArticle(id uint) bool {
	var count *int64
	global.DB().Model(model.Article{}).Where("id=?", id).Count(count)
	return *count >= 0
}
