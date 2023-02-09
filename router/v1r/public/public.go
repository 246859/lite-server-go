package public

import (
	"liteserver/router/route"
)

const (
	Public = "public"
)

type PublicRouter struct {
}

func (p PublicRouter) InitGroup() route.RouterGroupMap {
	return route.RouterGroupMap{
		Ping:           route.ApiGroup{Path: Ping, Router: PingRouter{}},
		Mail:           route.ApiGroup{Path: Mail, IsUrl: true, Router: MailRouter{}},
		User:           route.ApiGroup{Path: User, IsUrl: true, Router: UserSimpleRouter{}},
		Article:        route.ApiGroup{Path: Article, IsUrl: true, Router: ArticleRouter{}},
		Authentication: route.ApiGroup{Path: Authentication, IsUrl: true, Router: AuthenticationRouter{}},
	}
}
