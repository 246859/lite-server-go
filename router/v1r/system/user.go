package system

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
	"net/http"
)

const (
	User               = "user"
	UserInfo           = "info"
	UserList           = "list"
	UserUpdate         = "update"
	UserDelete         = "delete"
	UserLogout         = "logout"
	UserChangePassword = "changePassword"
)

type UserSystemRouter struct {
}

func (u UserSystemRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		UserInfo:           route.Api{Path: UserInfo, Method: http.MethodGet, Handler: v1r.UserController.UserInfo},
		UserList:           route.Api{Path: UserList, Method: http.MethodGet, Handler: v1r.UserController.UserList},
		UserUpdate:         route.Api{Path: UserUpdate, Method: http.MethodPost, Handler: v1r.UserController.UpdateUserInfo},
		UserDelete:         route.Api{Path: UserDelete, Method: http.MethodDelete, Handler: v1r.UserController.DeleteUser},
		UserLogout:         route.Api{Path: UserLogout, Method: http.MethodGet, Handler: v1r.UserController.Logout},
		UserChangePassword: route.Api{Path: UserChangePassword, Method: http.MethodGet, Handler: v1r.UserController.ChangePassword},
	}
}
