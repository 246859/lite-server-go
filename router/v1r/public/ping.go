package public

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

const (
	Ping = "ping"
)

type PingRouter struct {
}

func (p PingRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Ping: route.Api{Path: Ping, Method: http.MethodGet, Handler: v1r.Ping.Ping},
	}
}
