package mailutils

import (
	"github.com/246859/lite-server-go/config"
	"gopkg.in/gomail.v2"
)

const (
	TextHtml  = "text/html"
	TextPlain = "text/plain"
)

var goMailCfg *config.MailConfig

func SetConfig(mailConfig *config.MailConfig) {
	goMailCfg = mailConfig
}

func SendHtmlMail(subject, to, html string) error {
	message := newMessage(subject, to, TextHtml, html)
	return send(message)
}

func SendPlainMail(subject, to, plain string) error {
	message := newMessage(subject, to, TextPlain, plain)
	return send(message)
}

func newMessage(subject, to, contentType, body string) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", goMailCfg.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody(contentType, body)
	return message
}

func send(message *gomail.Message) error {
	dialer := gomail.NewDialer(goMailCfg.Host, goMailCfg.Port, goMailCfg.Username, goMailCfg.Password)
	return dialer.DialAndSend(message)
}
