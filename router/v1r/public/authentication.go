package public

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

// @Date: 2023-01-16 20:24:57
// 接口的路径
const (
	Authentication = "authen"
	Login          = "login"
	Register       = "register"
	ForgetPassword = "forgetPassword"
	RefreshToken   = "refreshToken"
)

type AuthenticationRouter struct {
}

func (a AuthenticationRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Login:          route.Api{Path: Login, Method: http.MethodGet, Handler: v1r.AutenController.Login},
		Register:       route.Api{Path: Register, Method: http.MethodPost, Handler: v1r.AutenController.Register},
		ForgetPassword: route.Api{Path: ForgetPassword, Method: http.MethodPost, Handler: v1r.AutenController.ForgetPassword},
		RefreshToken:   route.Api{Path: RefreshToken, Method: http.MethodGet, Handler: v1r.AutenController.RefreshToken},
	}
}
