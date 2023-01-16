package article

import "liteserver/router/route"

const (
	Comment = "comment"
)

type CommentRouter struct{}

func (c CommentRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
