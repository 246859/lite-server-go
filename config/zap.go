package config

import (
	"github.com/246859/lite-server-go/utils/fileutils"
	"io"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel   = "debug"
	InfoLevel    = "info"
	ErrorLevel   = "error"
	PanicLevel   = "panic"
	FatalLevel   = "fatal"
	WriteFile    = "file"
	WriteConsole = "console"
	WriteBoth    = "both"
)

// ZapConfig
// @Date: 2023-01-09 16:37:05
// @Description: zap日志配置结构体
type ZapConfig struct {
	Prefix     string         `yaml:"prefix" mapstructure:""prefix`
	TimeFormat string         `yaml:"timeFormat" mapstructure:"timeFormat"`
	Level      string         `yaml:"level" mapstructure:"level"`
	Caller     bool           `yaml:"caller" mapstructure:"caller"`
	StackTrace bool           `yaml:"stackTrace" mapstructure:"stackTrace"`
	Writer     string         `yaml:"writer" mapstructure:"writer"`
	Encode     string         `yaml:"encode" mapstructure:"encode"`
	LogFile    *LogFileConfig `yaml:"logFile" mapstructure:"logFile"`
}

// LogFileConfig
// @Date: 2023-01-09 16:38:45
// @Description: 日志文件配置结构体
type LogFileConfig struct {
	MaxSize  int      `yaml:"maxSize" mapstructure:"maxSize"`
	BackUps  int      `yaml:"backups" mapstructure:"backups"`
	Compress bool     `yaml:"compress" mapstructure:"compress"`
	Output   []string `yaml:"output" mapstructure:"output"`
}

// customTimeFormatEncoder
// @Date: 2023-01-09 19:39:16
// @Description: 自定义时间编码器
// @Param: t time.Time
// @Param: encoder zapcore.PrimitiveArrayEncoder
func customTimeFormatEncoder(cfg *ZapConfig) zapcore.TimeEncoder {
	return func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		builder := strings.Builder{}
		builder.WriteString(cfg.Prefix)
		builder.WriteString("\t")
		builder.WriteString(t.Format(cfg.TimeFormat))
		encoder.AppendString(builder.String())
	}
}

// BuildOptions
// @Date: 2023-01-09 20:46:39
// @Description: 构建Options
// @Receiver: z *ZapConfig
// @Param: levelEnabler zapcore.LevelEnabler
// @Return: options []zap.Option
func (z *ZapConfig) BuildOptions(levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if z.Caller {
		options = append(options, zap.AddCaller())
	}

	if z.StackTrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}

// ZapLevelEnabler
// @Date: 2023-01-09 16:03:23
// @Description:
// @Param: cfg *config.ZapConfig
// @Return: zapcore.LevelEnabler
func (z *ZapConfig) ZapLevelEnabler() zapcore.LevelEnabler {
	switch z.Level {
	case DebugLevel:
		return zap.DebugLevel
	case InfoLevel:
		return zap.InfoLevel
	case ErrorLevel:
		return zap.ErrorLevel
	case PanicLevel:
		return zap.PanicLevel
	case FatalLevel:
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}

// ZapConsoleWriterSyncer
// @Date: 2023-01-09 20:51:02
// @Description: 获取ConsoleWriterSyncer
// @Receiver: z *ZapConfig
// @Return: zapcore.WriteSyncer
func (z *ZapConfig) ZapConsoleWriterSyncer() zapcore.WriteSyncer {
	if z.Writer == WriteBoth || z.Writer == WriteConsole {
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.AddSync(io.Discard)
}

// ZapFileWriterSyncer
// @Date: 2023-01-09 20:52:02
// @Description: 获取FileWriterSyncer
// @Receiver: z *ZapConfig
// @Return: zapcore.WriteSyncer
func (z *ZapConfig) ZapFileWriterSyncer() zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if z.Writer == WriteBoth || z.Writer == WriteFile {
		// 添加日志输出器
		for _, path := range z.LogFile.Output {
			logger := &lumberjack.Logger{
				Filename:   fileutils.JoinPath(path),
				MaxSize:    z.LogFile.MaxSize,
				MaxBackups: z.LogFile.BackUps,
				Compress:   z.LogFile.Compress,
				LocalTime:  true,
			}
			syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
		}
	}
	return zap.CombineWriteSyncers(syncers...)
}

// ZapConsoleEncoder
// @Date: 2023-01-09 15:34:21
// @Description: 控制台编码器
// @Return zapcore.Encoder
func (z *ZapConfig) ZapConsoleEncoder() zapcore.Encoder {
	encodeConfig := z.newEncodeConfig()
	// 控制台开启颜色输出
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 简短的调用者输出
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encodeConfig.EncodeName = zapcore.FullNameEncoder

	return z.judgeConfigEncoder(encodeConfig)
}

// ZapFileEncoder
// @Date: 2023-01-09 20:59:35
// @Description: 文件编码器
// @Receiver: z *ZapConfig
// @Return: zapcore.Encoder
func (z *ZapConfig) ZapFileEncoder() zapcore.Encoder {
	encodeConfig := z.newEncodeConfig()
	// 文件关闭颜色输出
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 简短的调用者输出
	encodeConfig.EncodeCaller = zapcore.FullCallerEncoder
	// 完整的序列化logger名称
	encodeConfig.EncodeName = zapcore.FullNameEncoder

	return z.judgeConfigEncoder(encodeConfig)
}

// judgeConfigEncoder
// @Date: 2023-01-09 20:39:13
// @Description:
// @Param: config *config.ZapConfig
// @Param: encoderConfig zapcore.EncoderConfig
// @Return: zapcore.Encoder
func (z *ZapConfig) judgeConfigEncoder(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	// 最终的日志编码 json或者console
	switch z.Encode {
	case "json":
		{
			return zapcore.NewJSONEncoder(encoderConfig)
		}
	case "console":
		{
			return zapcore.NewConsoleEncoder(encoderConfig)
		}
	}
	// 默认console
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (z *ZapConfig) newEncodeConfig() zapcore.EncoderConfig {
	// 新建一个配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		FunctionKey:   "Func",
		LineEnding:    zapcore.DefaultLineEnding,
	}
	// 自定义时间格式
	encoderConfig.EncodeTime = customTimeFormatEncoder(z)
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	return encoderConfig
}
