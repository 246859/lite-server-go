package public

import (
	"liteserver/controller"
	"liteserver/router/route"
	"net/http"
)

// @Date: 2023-01-16 20:24:57
// 接口的路径
const (
	Authentication = "authen"
	Login          = "login"
	Register       = "register"
	Logout         = "logout"
	ForgetPassword = "forgetPassword"
	ChangePassword = "changePassword"
	RefreshToken   = "refreshToken"
)

var authenticationApi = controller.ControllerGroup.Public.Authentication

type AuthenticationRouter struct {
}

func (a AuthenticationRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Login:          &route.Api{Path: Login, Method: http.MethodGet, Handler: authenticationApi.Login},
		Register:       &route.Api{Path: Register, Method: http.MethodPost, Handler: authenticationApi.Register},
		Logout:         &route.Api{Path: Logout, Method: http.MethodGet, Handler: authenticationApi.Logout},
		ForgetPassword: &route.Api{Path: ForgetPassword, Method: http.MethodPost, Handler: authenticationApi.ForgetPassword},
		ChangePassword: &route.Api{Path: ChangePassword, Method: http.MethodPost, Handler: authenticationApi.ChangePassword},
		RefreshToken:   &route.Api{Path: RefreshToken, Method: http.MethodGet, Handler: authenticationApi.RefreshToken},
	}
}
