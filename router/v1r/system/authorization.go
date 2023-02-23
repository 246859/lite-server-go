package system

import "github.com/246859/lite-server-go/router/route"

const (
	Authorization = "auth"
)

type AuthorizationRouter struct {
}

func (a AuthorizationRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
