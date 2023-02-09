package service

import (
	"liteserver/service/article"
	"liteserver/service/mail"
	"liteserver/service/system"
)

var AppService = new(ServiceGroup)

type ServiceGroup struct {
	system.System
	mail.Mail
	article.Article
}
