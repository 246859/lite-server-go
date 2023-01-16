package system

import (
	"liteserver/router/route"
)

const (
	Authorization = "authorization"
)

type AuthorizationRouter struct {
}

func (a AuthorizationRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
