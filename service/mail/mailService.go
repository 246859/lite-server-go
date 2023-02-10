package mail

import (
	"go.uber.org/zap/buffer"
	"liteserver/global"
	"liteserver/model/sys/sysrep"
	"liteserver/resource"
	"liteserver/utils/mailutils"
	"liteserver/utils/uuidtool"
	"text/template"
)

type MailService struct{}

func (m *MailService) SendAuthMail(to string) (*sysrep.AuthMail, error) {
	// 生成UUID
	uuid := uuidtool.NewUUIDv5()
	// 截取后6位
	code := uuid[len(uuid)-6:]
	// 解析模板文件
	mailTemplate, err := template.ParseFS(resource.ResourceFS, "template/mail.html")
	if err != nil {
		return nil, err
	}
	// 创建writer
	buff := &buffer.Buffer{}
	// 创建邮件待解析数据
	mailData := &sysrep.AuthMail{Code: code, Expire: global.Config.MailConfig.Expire, To: to}
	// 解析数据
	err = mailTemplate.Execute(buff, mailData)
	if err != nil {
		return nil, err
	}
	err = mailutils.SimpleMailClient.SendHtmlMail(global.I18nRawCN("mail.authmail"), buff.String(), mailData.To)
	if err != nil {
		return nil, err
	}
	return mailData, nil
}
