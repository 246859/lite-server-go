package config

import (
	"github.com/spf13/viper"
)

// Config
// @Date: 2023-01-09 13:48:50
// @Description: 全局配置结构体，映射了整个配置文件
type Config struct {
	DataBaseConfig *DataBaseConfigGroup `mapstructure:"database"`
	RedisConfig    *RedisConfig         `mapstructure:"redis"`
	JwtConfig      *JwtConfig           `mapstructure:"jwt"`
	ServerConfig   *ServerConfig        `mapstructure:"server"`
	ZapConfig      *ZapConfig           `mapstructure:"zap"`
	I18nConfig     *I18nConfig          `mapstructure:"i18n"`
	MailConfig     *MailConfig          `mapstructure:"mail"`
}

// ReadConfig
// @Date: 2023-01-08 21:55:36
// @Description: 读取应用配置
// @Param: cfgPath string
// @Return: *viper.Viper
// @Return: error
func ReadConfig(cfgPath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(cfgPath)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

// RefreshConfig
// @Date: 2023-01-08 21:55:22
// @Description: 刷新应用配置
// @Param: viper viper.Viper
// @Return: *Config
// @Return: error
func RefreshConfig(viper *viper.Viper) (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
