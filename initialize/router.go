package initialize

import (
	"github.com/gin-gonic/gin"
	"liteserver/router"
	"liteserver/router/route"
	"liteserver/utils/ginutils"
)

func InitRouter(engine *gin.Engine) {
	group := engine.RouterGroup
	httpRouterGroup := group.Group(router.GinRouter.Version)
	registerRouterGroup(httpRouterGroup, router.GinRouter.SystemRoute)
	registerRouterGroup(httpRouterGroup, router.GinRouter.PublicRoute)
}

// registerRouter
// @Date 2023-01-16 20:20:31
// @Param group gin.RouterGroup
// @Param routeMap route.RouterMap
// @Description: 注册接口路由
func registerRouter(httpRouterGroup *gin.RouterGroup, route route.Router) {
	for k, api := range route.InitRouter() {
		if len(k) > 0 { // 如果接口路径为空，则跳过
			httpMethod := ginutils.JudgeMethod(api.Method, httpRouterGroup)
			if httpMethod != nil { // 如果http方法不存在，则跳过
				httpMethod(api.Path, append(api.ConfigMiddleWare(), api.Handler)...)
			}
		}
	}
}

// registerRouterGroup
// @Date 2023-01-16 21:35:04
// @Param httpRouterGroup *gin.RouterGroup
// @Param routerGroup route.RouterGroup
// @Description: 注册路由分组
func registerRouterGroup(httpRouterGroup *gin.RouterGroup, routerGroup route.RouterGroup) {
	for path, group := range routerGroup.InitGroup() {
		if group.Group != nil {
			registerRouterGroup(httpRouterGroup.Group(path, group.ConfigMiddleWare()...), group.Group)
		} else if group.Router != nil {
			registerRouter(httpRouterGroup, group.Router)
		}
	}
}
