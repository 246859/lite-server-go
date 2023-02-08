package config

// MailConfig
// @Date 2023-02-08 15:32:32
// @Description: 邮件客户端连接配置
type MailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	// Expire
	// @Date 2023-02-08 17:20:50
	// @Description: 验证码过期时间
	Expire int `mapstructure:"expire"`
	// Limit
	// @Date 2023-02-08 17:20:56
	// @Description: 邮件重复发送限制
	Limit int `mapstructure:"limit"`
}
