package public

import (
	"github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model/request"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ArticleController
// @Date 2023-02-09 19:59:39
// @Description: 文章公共接口
type ArticleController struct {
}

// ArticleInfo
// @Date 2023-02-20 17:57:40
// @Param ctx *gin.Context
// @Method http.MethodGet
// @Url /article/info
// @Description: 查询单个文章信息
func (a ArticleController) ArticleInfo(ctx *gin.Context) {
	articleId := ctx.Query("articleId")
	ints, err := strconv.Atoi(articleId)
	if err != nil || ints < 0 {
		responseuils.FailWithMsg(ctx, global.I18nRawCN("article.invalidId"))
		return
	}

	if article, err := v1c.ArticleService.Article(ints); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, article, global.I18nRawCN("article.query.ok"))
	}
}

// ArticleList
// @Date 2023-02-20 17:58:35
// @Param ctx *gin.Context
// @Method http.MethodGet
// @Url /article/list
// @Description: 分页查询
func (a ArticleController) ArticleList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	// 参数解析
	if err := ctx.ShouldBind(&pageInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	if page, err := v1c.ArticleService.ArticlePage(pageInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, page, global.I18nRawCN("article.query.ok"))
	}
}

// ArticleComment
// @Date 2023-02-24 19:52:27
// @Method http.MethodGet
// @Url /article/comment
// @Param articleId string 文章ID
// @Param page int 页数
// @Param size int 页容量
// @Param desc bool 顺序
// @Description: 获取文章的评论列表
func (ArticleController) ArticleComment(ctx *gin.Context) {

}
