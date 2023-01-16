package system

import (
	"liteserver/router/route"
)

type SystemRouter struct {
}

func (s SystemRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		Authorization:  &route.ApiGroup{Path: Authorization, Router: AuthorizationRouter{}},
		Authentication: &route.ApiGroup{Path: Authentication, Router: AuthenticationRouter{}},
	}
}
