package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfig struct {
	Level         zapcore.Level `yaml:"level" mapstructure:"level"`
	Development   bool          `yaml:"development" mapstructure:"development"`
	Caller        bool          `yaml:"caller" mapstructure:"caller"`
	Encode        string        `yaml:"encode" mapstructure:"encode"`
	OutPutPath    []string      `yaml:"outputPath" mapstructure:"outputPath"`
	ErrOutPutPath []string      `yaml:"errOutPutPath" mapstructure:"errOutputPath"`
}

// GetConfig
// @Date: 2023-01-09 14:46:45
// @Description: 生成Zap配置
// @Receiver: z *ZapConfig
// @Return: *zap.Config
func (z *ZapConfig) GetConfig() *zap.Config {

	config := &zap.Config{
		Level:            zap.NewAtomicLevel(),
		Development:      z.Development,
		DisableCaller:    z.Caller,
		Encoding:         z.Encode,
		OutputPaths:      z.OutPutPath,
		ErrorOutputPaths: z.ErrOutPutPath,
	}
	switch z.Level {
	case zap.DebugLevel:
		{
			config.Level.SetLevel(zap.DebugLevel)
		}
	case zap.InfoLevel:
		{
			config.Level.SetLevel(zap.InfoLevel)
		}
	case zap.ErrorLevel:
		{
			config.Level.SetLevel(zap.ErrorLevel)
		}
	case zap.PanicLevel:
		{
			config.Level.SetLevel(zap.PanicLevel)
		}
	default:
		{
			config.Level.SetLevel(zap.InfoLevel)
		}
	}
	return config
}
