package article

import "liteserver/router/route"

const (
	Article = "article"
)

type ArticleRouterGroup struct{}

func (a ArticleRouterGroup) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		Info:    &route.ApiGroup{Path: Info, Router: ArticleInfoRouter{}},
		Comment: &route.ApiGroup{Path: Comment, Router: CommentRouter{}},
	}
}
