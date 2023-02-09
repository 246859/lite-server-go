package public

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

const (
	Article     = "article"
	ArticleInfo = "info"
	ArticleList = "list"
)

type ArticleRouter struct {
}

func (a ArticleRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		ArticleInfo: route.Api{Path: ArticleInfo, Method: http.MethodGet, Handler: v1r.ArticleController.ArticleInfo},
		ArticleList: route.Api{Path: ArticleList, Method: http.MethodGet, Handler: v1r.ArticleController.ArticleList},
	}
}
