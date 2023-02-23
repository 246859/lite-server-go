package system

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
	"net/http"
)

const (
	Article = "article"

	ArticleDelete = "delete"
	ArticleUpdate = "update"

	ArticleCreate = "create"
)

type ArticleSystemRouter struct {
}

func (a ArticleSystemRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		ArticleDelete: route.Api{Path: ArticleDelete, Method: http.MethodDelete, Handler: v1r.ArticleModifyController.DeleteArticle},
		ArticleUpdate: route.Api{Path: ArticleUpdate, Method: http.MethodPost, Handler: v1r.ArticleModifyController.UpdateArticle},
		ArticleCreate: route.Api{Path: ArticleCreate, Method: http.MethodPost, Handler: v1r.ArticleModifyController.CreateArticle},
	}
}
