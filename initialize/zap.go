package initialize

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"liteserver/config"
)

// InitZap
// @Date: 2023-01-09 14:52:27
// @Description: Zap初始化
// @Param: config *config.ZapConfig
// @Return: *zap.Logger
func InitZap(config *config.ZapConfig) *zap.Logger {
	encoder := zapEncoder(config)
	levelEnabler := zapLevelEnabler(config)
	subCore, options := tee(config, encoder, levelEnabler)
	return zap.New(subCore, options...)
}

// tee
// @Date: 2023-01-09 16:48:01
// @Description: 将所有的zapcore 结合在一起
// @Param: cfg *config.ZapConfig
// @Param: encoder zapcore.Encoder
// @Param: levelEnabler zapcore.LevelEnabler
// @Return: []zapcore.Core
func tee(cfg *config.ZapConfig, encoder zapcore.Encoder, levelEnabler zapcore.LevelEnabler) (core zapcore.Core, options []zap.Option) {
	sink := zapWriteSyncer(cfg)
	return zapcore.NewCore(encoder, sink, levelEnabler), buildOptions(cfg, levelEnabler)
}

// buildOptions
// @Date: 2023-01-09 17:19:27
// @Description: 构建Core配置选项
// @Param: cfg *config.ZapConfig
// @Param: levelEnabler zapcore.LevelEnabler
// @Param: errsink *zapcore.WriteSyncer
// @Return: options []zap.Option
func buildOptions(cfg *config.ZapConfig, levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if cfg.Caller {
		options = append(options, zap.AddCaller())
	}

	if cfg.StackTrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}

// encodeConfig
// @Date: 2023-01-09 15:34:21
// @Description: 根据配置自定义Zap的编码器，也就是日志的输出样式定义
// @Return zapcore.Encoder 日志编码器
func zapEncoder(config *config.ZapConfig) zapcore.Encoder {
	// 新建一个配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
	}
	// ISO时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 简短的调用者输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	// 最终的日志编码 json或者console
	switch config.Encode {
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

// zapWriteSyncer
// @Date: 2023-01-09 15:57:04
// @Description: 日志书写器，负责将日志写入到什么地方，默认设置文件和控制台
// @Return: zapcore.WriteSyncer 日志书写器
func zapWriteSyncer(cfg *config.ZapConfig) (outsink zapcore.WriteSyncer) {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志控制台输出，就加入控制台书写器
	if cfg.Writer == config.WriteBoth || cfg.Writer == config.WriteConsole {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if cfg.Writer == config.WriteBoth || cfg.Writer == config.WriteFile {
		// 添加日志输出器
		for _, path := range cfg.LogFile.Output {
			logger := &lumberjack.Logger{
				Filename:   path,
				MaxSize:    cfg.LogFile.MaxSize,
				MaxBackups: cfg.LogFile.BackUps,
				Compress:   cfg.LogFile.Compress,
				LocalTime:  true,
			}
			syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
		}
	}
	return zap.CombineWriteSyncers(syncers...)
}

// zapLevelEnabler
// @Date: 2023-01-09 16:03:23
// @Description:
// @Param: cfg *config.ZapConfig
// @Return: zapcore.LevelEnabler
func zapLevelEnabler(cfg *config.ZapConfig) zapcore.LevelEnabler {
	switch cfg.Level {
	case config.DebugLevel:
		return zap.DebugLevel
	case config.InfoLevel:
		return zap.InfoLevel
	case config.ErrorLevel:
		return zap.ErrorLevel
	case config.PanicLevel:
		return zap.PanicLevel
	case config.FatalLevel:
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}
