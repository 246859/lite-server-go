package system

import (
	"github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/model"
	"github.com/246859/lite-server-go/utils/jwtutils"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ArticleModifyController
// @Date 2023-02-09 20:05:03
// @Description: 文章修改相关接口
type ArticleModifyController struct {
}

// DeleteArticle
// @Date 2023-02-09 20:05:50
// @Param ctx *gin.Context
// @Method http.MethodDelete
// @Description: 删除文章接口
func (a ArticleModifyController) DeleteArticle(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	ints, err := strconv.Atoi(articleId)
	if err != nil || ints < 0 {
		responseuils.FailWithMsg(ctx, "非法的文章ID")
		return
	}

	if err := v1c.ArticleService.DeleteArticle(ints); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsg(ctx, "文章删除成功")
	}
}

// UpdateArticle
// @Date 2023-02-20 18:21:48
// @Param ctx *gin.Context
// @Method http.MethodPost
// @Description: 更新文章接口
func (a ArticleModifyController) UpdateArticle(ctx *gin.Context) {
	var articleInfo model.Article
	if err := ctx.ShouldBindUri(&articleInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	if err := v1c.ArticleService.UpdateArticle(&articleInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsg(ctx, "文章更新成功")
	}
}

// CreateArticle
// @Date 2023-02-20 18:25:25
// @Param ctx *gin.Context
// @Method http.MethodPost
// @Description: 创建文章接口
func (a ArticleModifyController) CreateArticle(ctx *gin.Context) {
	var articleInfo model.Article
	if err := ctx.ShouldBind(&articleInfo); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	if err := v1c.ArticleService.CreateArticle(&articleInfo, claims.UserClaims); err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsg(ctx, "文章创建成功")
	}
}
