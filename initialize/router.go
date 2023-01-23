package initialize

import (
	"github.com/gin-gonic/gin"
	"liteserver/config"
	"liteserver/router"
	"liteserver/router/route"
	"liteserver/utils/ginutils"
	"path/filepath"
)

func InitRouter(engine *gin.Engine, cfg *config.ServerConfig) {
	// 公共静态文件映射
	engine.StaticFS("/static", gin.Dir(filepath.Join(cfg.WorkDir, cfg.StaticDir), true))
	// 获取版本号
	httpRouterGroup := engine.Group(router.GinRouter.Version)
	// 注册路由
	registerRouterGroup(httpRouterGroup, router.GinRouter)

}

// registerRouter
// @Date 2023-01-16 20:20:31
// @Param group gin.RouterGroup
// @Param routeMap route.RouterMap
// @Description: 注册接口路由
func registerRouter(httpRouterGroup *gin.RouterGroup, r route.Router) {
	for k, api := range r.InitRouter() {
		if len(k) > 0 { // 如果接口路径为空，则跳过
			httpMethod := ginutils.JudgeMethod(api.Method, httpRouterGroup)
			if httpMethod != nil { // 如果http方法不存在，则跳过
				var mds route.Md
				if api.Mds != nil { // 如果中间件列表不为空
					mds = api.Mds.ConfigMiddleWare()
				}
				// handler就是接口，永远放在中间件后面执行
				httpMethod(api.Path, append(mds, api.Handler)...)
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
			// gin路由组创建，如果路由组的path不为空，那么就需要额外创建一个gin路由组
			var parent *gin.RouterGroup
			if len(group.Path) > 0 {
				parent = httpRouterGroup.Group(path)
			} else {
				parent = httpRouterGroup
			}

			var mds route.Md // 路由组中间件加载
			if group.Mds != nil {
				mds = group.Mds.ConfigMiddleWare()
			}

			parent.Use(mds...)
			// 进入下一层
			registerRouterGroup(parent, group.Group)
		} else if group.Router != nil {
			registerRouter(httpRouterGroup, group.Router)
		}
	}
}
