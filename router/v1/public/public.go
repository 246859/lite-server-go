package public

import (
	"liteserver/router/route"
	"liteserver/router/v1/public/article"
)

type PublicRouter struct {
}

func (p PublicRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		article.Article: &route.ApiGroup{Path: article.Article, Group: article.ArticleRouterGroup{}},
		Ping:            &route.ApiGroup{Path: Ping, Router: PingRouter{}},
	}
}
