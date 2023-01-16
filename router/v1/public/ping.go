package public

import (
	"liteserver/controller"
	"liteserver/router/route"
	"net/http"
)

const (
	Ping = "ping"
)

var pingApi = controller.ControllerGroup.Public.Ping

type PingRouter struct {
}

func (p PingRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Ping: &route.Api{Path: Ping, Method: http.MethodGet, Handler: pingApi.Ping},
	}
}
