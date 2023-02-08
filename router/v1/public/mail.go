package public

import (
	"liteserver/controller"
	"liteserver/router/route"
	"net/http"
)

const (
	Mail     = "mail"
	AuthMail = "authMail"
)

var mailApi = controller.ControllerGroup.Public.Mail

type MailRouter struct {
}

func (m MailRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		AuthMail: &route.Api{Path: AuthMail, Method: http.MethodGet, Handler: mailApi.SendAuthMail},
	}
}
