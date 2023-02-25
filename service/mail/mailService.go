package mail

import (
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/resource"
	"github.com/246859/lite-server-go/utils/mailutils"
	"github.com/246859/lite-server-go/utils/uuidtool"
	"go.uber.org/zap/buffer"
	"text/template"
)

type MailService struct{}

func (m *MailService) SendAuthMail(to string) (*response.AuthMail, error) {
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
	mailData := &response.AuthMail{Code: code, Expire: global.Config.MailConfig.Expire, To: to}
	// 解析数据
	err = mailTemplate.Execute(buff, mailData)
	if err != nil {
		return nil, err
	}
	err = mailutils.SendHtmlMail(global.I18nRawCN("mail.authmail"), mailData.To, buff.String())
	if err != nil {
		return nil, err
	}
	return mailData, nil
}
