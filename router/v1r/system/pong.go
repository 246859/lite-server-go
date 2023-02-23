package system

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
	"net/http"
)

const (
	Pong = "pong"
)

type PongRouter struct {
}

func (p PongRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Pong: route.Api{Path: Pong, Method: http.MethodGet, Handler: v1r.Pong.Pong},
	}
}
