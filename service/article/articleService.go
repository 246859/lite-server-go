package article

import (
	"errors"
	"liteserver/global"
	"liteserver/model/article"
	"liteserver/model/sys/sysreq"
	"liteserver/service/page"
	"liteserver/utils/jwtutils"
)

type ArticleService struct {
}

// Article
// @Date 2023-02-20 16:25:22
// @Param articleId string
// @Return *article.Article
// @Return error
// @Description: 根据单个id查询文章内容
func (a ArticleService) Article(articleId int) (*article.Article, error) {
	var articleinfo article.Article
	err := global.DB().Model(articleinfo).Where("id=?", articleId).First(&articleinfo).Error
	if err != nil {
		return nil, err
	} else {
		return &articleinfo, nil
	}
}

// ArticlePage
// @Date 2023-02-20 17:36:17
// @Param pageInfo sysreq.PageInfo
// @Return []article.HeadInfo
// @Return error
// @Description: 分页查询
func (a ArticleService) ArticlePage(pageInfo sysreq.PageInfo) ([]article.HeadInfo, error) {
	page := new(page.PageService).SelectPage(pageInfo)
	var articleList []article.HeadInfo
	if err := page(global.Model(article.Article{}), nil, &articleList).Error; err != nil {
		return nil, err
	} else {
		return articleList, nil
	}
}

// CreateArticle
// @Date 2023-02-20 16:46:19
// @Param article *article.Article
// @Return error
// @Description: 创建一篇新文章
func (a ArticleService) CreateArticle(newArticle *article.Article, claims jwtutils.UserClaims) error {
	newArticle.UserId = claims.UserId
	err := global.DB().Model(article.Article{}).Create(newArticle).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteArticle
// @Date 2023-02-20 16:51:16
// @Param id string
// @Return error
// @Description: 删除文章
func (a ArticleService) DeleteArticle(id int) error {
	err := global.DB().Model(article.Article{}).Delete("id=?", id).Error
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
func (a ArticleService) UpdateArticle(articleOld *article.Article) error {
	if !a.HasArticle(articleOld.ID) {
		return errors.New("文章不存在")
	}
	err := global.Model(article.Article{}).Updates(articleOld).Error
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
	global.DB().Model(article.Article{}).Where("id=?", id).Count(count)
	return *count >= 0
}
