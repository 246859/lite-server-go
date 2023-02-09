package public

import "liteserver/router/route"

const (
	User = "user"
)

type UserSimpleRouter struct {
}

func (u UserSimpleRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
