package public

import (
	"github.com/gin-gonic/gin"
	"liteserver/controller/v1c"
	"liteserver/model/sys/sysreq"
	"liteserver/utils/response"
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
// @Description: 查询单个文章信息
func (a ArticleController) ArticleInfo(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	ints, err := strconv.Atoi(articleId)
	if err != nil || ints < 0 {
		response.FailWithMsg(ctx, "非法的文章ID")
		return
	}

	if article, err := v1c.ArticleService.Article(ints); err != nil {
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithMsgAndData(ctx, article, "文章信息查询成功")
	}
}

// ArticleList
// @Date 2023-02-20 17:58:35
// @Param ctx *gin.Context
// @Method http.MethodGet
// @Description: 分页查询
func (a ArticleController) ArticleList(ctx *gin.Context) {
	var pageInfo sysreq.PageInfo
	// 参数解析
	if err := ctx.ShouldBind(&pageInfo); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	if page, err := v1c.ArticleService.ArticlePage(pageInfo); err != nil {
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithMsgAndData(ctx, page, "文章信息查询成功")
	}
}
