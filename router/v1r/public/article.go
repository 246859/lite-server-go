package public

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
	"net/http"
)

const (
	Article     = "article"
	ArticleInfo = "info"
	ArticleList = "list"

	ArticleComment = "comment"
)

type ArticleRouter struct {
}

func (a ArticleRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		ArticleInfo:    route.Api{Path: ArticleInfo, Method: http.MethodGet, Handler: v1r.ArticleController.ArticleInfo},
		ArticleList:    route.Api{Path: ArticleList, Method: http.MethodGet, Handler: v1r.ArticleController.ArticleList},
		ArticleComment: route.Api{Path: ArticleComment, Method: http.MethodGet, Handler: v1r.ArticleController.ArticleComment},
	}
}
