package config

// I18nConfig
// @Date 2023-01-12 15:15:54
// @Description: i18n配置文件结构体
type I18nConfig struct {
	Dir    string `mapstructure:"dir"`
	Suffix string `mapstructure:"suffix"`
}
