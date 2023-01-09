package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DataBaseConfig *DataBaseConfigGroup `mapstructure:"database"`
	RedisConfig    *RedisConfig         `mapstructure:"redis"`
	JwtConfig      *JwtConfig           `mapstructure:"jwt"`
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
	fmt.Println(viper.AllSettings())
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
