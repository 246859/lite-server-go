package router

import (
	"liteserver/router/route"
	"liteserver/router/v1/public"
	"liteserver/router/v1/system"
)

// GinRouter
// @Date: 2023-01-23 21:27:34
// 应用总路由
var GinRouter = AppRouter{Version: "v1"}

type AppRouter struct {
	Version string
}

func (a AppRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		system.System: &route.ApiGroup{ // System部分路由
			Mds:   route.GeneralMiddleware(),
			Group: system.SystemRouter{},
		},
		public.Public: &route.ApiGroup{ // Public部分路由
			Mds:   route.GeneralMiddleware(),
			Group: public.PublicRouter{},
		},
	}
}
