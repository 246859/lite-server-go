package system

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
	"net/http"
)

const (
	User               = "user"
	UserSimpleInfo     = "simple"
	UserBasicInfo      = "basic"
	UserSimpleInfoList = "simples"
	UserBasicInfoList  = "basics"
	UserUpdate         = "update"
	UserDelete         = "delete"
	UserLogout         = "logout"
	UserChangePassword = "changePassword"
)

type UserSystemRouter struct {
}

func (u UserSystemRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		UserSimpleInfo:     route.Api{Path: UserSimpleInfo, Method: http.MethodGet, Handler: v1r.UserController.UserSimpleInfo},
		UserBasicInfo:      route.Api{Path: UserBasicInfo, Method: http.MethodGet, Handler: v1r.UserController.UserBasicInfo},
		UserSimpleInfoList: route.Api{Path: UserSimpleInfoList, Method: http.MethodGet, Handler: v1r.UserController.ListUserSimpleInfo},
		UserBasicInfoList:  route.Api{Path: UserBasicInfoList, Method: http.MethodGet, Handler: v1r.UserController.ListUserBasicInfo},
		UserUpdate:         route.Api{Path: UserUpdate, Method: http.MethodPost, Handler: v1r.UserController.UpdateUserInfo},
		UserDelete:         route.Api{Path: UserDelete, Method: http.MethodDelete, Handler: v1r.UserController.DeleteUser},
		UserLogout:         route.Api{Path: UserLogout, Method: http.MethodDelete, Handler: v1r.UserController.Logout},
		UserChangePassword: route.Api{Path: UserChangePassword, Method: http.MethodGet, Handler: v1r.UserController.ChangePassword},
	}
}
