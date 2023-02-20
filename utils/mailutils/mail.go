package mailutils

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"liteserver/config"
	"time"
)

var SimpleMailClient *MailClient

func NewMailClient(mailConfig *config.MailConfig) *MailClient {
	return &MailClient{MailCfg: mailConfig}
}

type MailClient struct {
	MailCfg    *config.MailConfig
	smtpClient *mail.SMTPClient
}

// SetConfig
// @Date 2023-02-08 16:01:31
// @Param cfg *config.MailConfig
// @Description: 设置客户端配置
func (m *MailClient) SetConfig(cfg *config.MailConfig) {
	m.MailCfg = cfg
}

func (m *MailClient) SendHtmlMail(subject string, body string, to ...string) error {
	email := mail.NewMSG()
	email.SetFrom(m.MailCfg.Username).
		AddTo(to...).
		SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	return m.send(email)
}

// send
// @Date 2023-02-08 16:04:05
// @Param email *mail.Email
// @Return error
// @Description: 发送邮件
func (m *MailClient) send(email *mail.Email) error {
	if email.Error != nil {
		return email.Error
	}
	if err := m.connect(); err != nil {
		return err
	}
	if err := email.Send(m.smtpClient); err != nil {
		return err
	}
	return nil
}

// Connect
// @Date 2023-02-08 15:47:43
// @Return error
// @Description: 与邮件服务器建立链接
func (m *MailClient) connect() error {
	server := mail.NewSMTPClient()
	server.Host = m.MailCfg.Host
	server.Port = m.MailCfg.Port
	server.Username = m.MailCfg.Username
	server.Password = m.MailCfg.Password
	server.ConnectTimeout = time.Second * 20
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = true

	client, err := server.Connect()
	if err != nil {
		return err
	} else {
		m.smtpClient = client
		return nil
	}
}
