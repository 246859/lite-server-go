package test

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestSendMail(t *testing.T) {
	dialer := gomail.NewDialer("smtp.qq.com", 465, "2633565580@qq.com", "uulvlwddcsyjdide")
	message := gomail.NewMessage()
	message.SetHeader("From", "2633565580@qq.com")
	message.SetHeader("To", "2633565580@qq.com")
	message.SetBody("text/plain", "hello world")
	fmt.Println(dialer.DialAndSend(message))
}

func TestName(t *testing.T) {
	ch := make(chan *gomail.Message)

	message := gomail.NewMessage()
	message.SetHeader("From", "2633565580@qq.com")
	message.SetHeader("To", "2633565580@qq.com")
	message.SetBody("text/plain", "hello world")
	ch <- message
	close(ch)
}
