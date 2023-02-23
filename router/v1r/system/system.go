package system

import "github.com/246859/lite-server-go/router/route"

const (
	System = "system"
)

type SystemRouter struct {
}

func (s SystemRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		Pong:          route.ApiGroup{Path: Pong, Router: PongRouter{}},
		User:          route.ApiGroup{Path: User, IsUrl: true, Router: UserSystemRouter{}},
		Article:       route.ApiGroup{Path: Article, IsUrl: true, Router: ArticleSystemRouter{}},
		Authorization: route.ApiGroup{Path: Authorization, IsUrl: true, Router: AuthorizationRouter{}},
		File:          route.ApiGroup{Path: File, IsUrl: true, Router: FileRouter{}},
	}
}
