package pubrep

// AuthMail
// @Date 2023-02-08 17:10:25
// @Description: 验证码邮件结构体
type AuthMail struct {
	To     string
	Code   string
	Expire int
}
