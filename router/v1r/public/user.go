package public

import "github.com/246859/lite-server-go/router/route"

const (
	User = "user"
)

type UserSimpleRouter struct {
}

func (u UserSimpleRouter) InitRouter() route.RouterMap {
	return route.RouterMap{}
}
