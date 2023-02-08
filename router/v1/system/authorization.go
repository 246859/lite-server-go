package system

import (
	"liteserver/router/route"
)

const (
	Authorization = "auth"
)

type AuthorizationRouter struct {
}

func (a AuthorizationRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
