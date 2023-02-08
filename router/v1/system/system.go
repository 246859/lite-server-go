package system

import (
	"liteserver/router/route"
)

const (
	System = "system"
)

type SystemRouter struct {
}

func (s SystemRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		Authorization: &route.ApiGroup{Path: Authorization, IsUrl: true, Router: AuthorizationRouter{}},
	}
}
