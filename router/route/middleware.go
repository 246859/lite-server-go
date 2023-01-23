package route

import (
	"github.com/gin-gonic/gin"
	"liteserver/middleware"
)

const (
	// RouterCfgKey 路由配置的key
	RouterCfgKey = "cfg"
)

// Md 别名 中间件
type Md = []gin.HandlerFunc

// MiddleWareRoute
// @Date 2023-01-16 21:36:55
// @Description: 中间件和对应的配置
type MiddleWareRoute struct {
	// 中间件
	Mds Md
	// 配置，在请求到达时会将对应的配置传入context
	Cfg *Config
}

// Config
// @Date 2023-01-23 19:54:22
// @Description: 路由配置结构体
type Config struct {
	// 是否需要认证
	Auth bool
	// 访问次数控制
	Limit int
	// 访问权限控制
	// TODO 暂时没想好perm的类型
	Perm interface{}
}

// ConfigMiddleWare
// @Date 2023-01-16 21:27:54
// @Return []gin.HandlerFunc
// @Description: 路由中间件配置
func (m MiddleWareRoute) ConfigMiddleWare() []gin.HandlerFunc {
	config := func(c *gin.Context) {
		c.Set(RouterCfgKey, m.Cfg)
		c.Next()
	}
	return append([]gin.HandlerFunc{config}, m.Mds...)
}

// GeneralMiddleware
// @Date 2023-01-23 21:04:50
// @Return *MiddleWareRoute
// @Method
// @Description: 同游路由中间件配置
func GeneralMiddleware() *MiddleWareRoute {
	md := &MiddleWareRoute{
		Mds: Md{middleware.JwtMiddleWare()},
	}
	return md
}
