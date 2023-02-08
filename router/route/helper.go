package route

import (
	"errors"
	"liteserver/middleware"
)

// IsNeedToAuth
// @Date 2023-02-06 20:09:13
// @Param config any
// @Return bool
// @Return error
// @Method
// @Description: 是否需要JWT认证
func IsNeedToAuth(config any) (bool, error) {
	if cfg, ok := config.(Config); ok {
		return cfg.Auth, nil
	}
	return false, errors.New("不是路由配置类型")
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
