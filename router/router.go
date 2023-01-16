package router

import (
	"liteserver/router/v1/public"
	"liteserver/router/v1/system"
)

var GinRouter = AppRouter{Version: "v1"}

type AppRouter struct {
	Version string
	// 后台接口分组
	SystemRoute system.SystemRouter
	// 前台接口分组
	PublicRoute public.PublicRouter
}
