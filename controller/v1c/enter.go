package v1c

import (
	"github.com/246859/lite-server-go/service"
)

var (
	Service = service.AppService
)

var (
	MailSevice     = Service.Mail
	SystemService  = Service.System
	ArticleService = Service.Article
	UserService    = Service.UserService

	FileService = Service.File
)
