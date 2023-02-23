package response

// AuthMail
// @Date 2023-02-08 17:10:25
// @Description: 验证码邮件展示信息
type AuthMail struct {
	To     string `json:"to"`
	Code   string `json:"code"`
	Expire int    `json:"expire"`
}
