package public

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
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
