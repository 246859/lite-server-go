package system

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

const (
	Article = "article"

	ArticleDelete = "delete"
	ArticleUpdate = "update"
)

type ArticleSystemRouter struct {
}

func (a ArticleSystemRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		ArticleDelete: route.Api{Path: ArticleDelete, Method: http.MethodDelete, Handler: v1r.ArticleModifyController.DeleteArticle},
		ArticleUpdate: route.Api{Path: ArticleUpdate, Method: http.MethodPost, Handler: v1r.ArticleModifyController.UpdateArticle},
	}
}
