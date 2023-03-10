package router

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r/public"
	"github.com/246859/lite-server-go/router/v1r/system"
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
		system.System: route.ApiGroup{ // Private部分路由 需要jwt验证
			Mds:   route.GeneralMiddleware(),
			Group: system.SystemRouter{},
			IsUrl: true,
		},
		public.Public: route.ApiGroup{ // Public部分路由 不需要jwt验证
			Group: public.PublicRouter{},
			IsUrl: true,
		},
	}
}
