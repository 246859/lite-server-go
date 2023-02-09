package v1c

import "liteserver/service"

var (
	Service = service.AppService
)

var (
	MailSevice     = Service.Mail
	SystemService  = Service.System
	ArticleService = Service.Article
)
