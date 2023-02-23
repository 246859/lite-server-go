package public

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
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
