package service

import (
	"github.com/246859/lite-server-go/service/article"
	"github.com/246859/lite-server-go/service/file"
	"github.com/246859/lite-server-go/service/mail"
	"github.com/246859/lite-server-go/service/system"
)

var AppService = new(ServiceGroup)

type ServiceGroup struct {
	system.System
	mail.Mail
	article.Article
	file.File
}
