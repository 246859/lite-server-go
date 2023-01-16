package route

import "github.com/gin-gonic/gin"

const (
	RouterCfgKey = "apiCfg"
)

// MiddleWareRoute
// @Date 2023-01-16 21:36:55
// @Description: 中间件和对应的配置
type MiddleWareRoute struct {
	// 中间件
	Middleware []gin.HandlerFunc
	// 配置，在请求到达时会将对应的配置传入context
	Config gin.H
}

// ConfigMiddleWare
// @Date 2023-01-16 21:27:54
// @Return []gin.HandlerFunc
// @Description: 路由中间件配置
func (m MiddleWareRoute) ConfigMiddleWare() []gin.HandlerFunc {
	config := func(c *gin.Context) {
		c.Set(RouterCfgKey, m.Config)
		c.Next()
	}
	return append([]gin.HandlerFunc{config}, m.Middleware...)
}
