package initialize

import (
	"liteserver/config"
	"liteserver/utils/mailutils"
)

// InitMail
// @Date 2023-02-08 15:56:56
// @Param config *config.MailConfig
// @Method
// @Description: 初始化邮件服务器
func InitMail(config *config.MailConfig) {
	mailClient := mailutils.NewMailClient(config)
	mailutils.SimpleMailClient = mailClient
}
