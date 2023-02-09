package public

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

const (
	Mail     = "mail"
	AuthMail = "authMail"
)

type MailRouter struct {
}

func (m MailRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		AuthMail: route.Api{Path: AuthMail, Method: http.MethodGet, Handler: v1r.MailController.SendAuthMail},
	}
}
