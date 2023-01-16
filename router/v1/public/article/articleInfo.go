package article

import "liteserver/router/route"

const (
	Info = "info"
)

type ArticleInfoRouter struct{}

func (a ArticleInfoRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
