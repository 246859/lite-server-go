package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"liteserver/config"
	"liteserver/utils/fileutils"
)

// InitZap
// @Date: 2023-01-09 14:52:27
// @Description: Zap初始化
// @Param: config *config.ZapConfig
// @Return: *zap.Logger
func InitZap(cfg *config.ZapConfig) *zap.Logger {
	initZapOutput(cfg.LogFile.Output)
	subCore, options := tee(cfg)
	logger := zap.New(subCore, options...)
	zap.ReplaceGlobals(logger)
	return logger
}

// tee
// @Date: 2023-01-09 16:48:01
// @Description: 将所有的zapcore 结合在一起
// @Param: cfg *config.ZapConfig
// @Param: encoder zapcore.Encoder
// @Param: levelEnabler zapcore.LevelEnabler
// @Return: []zapcore.Core
func tee(cfg *config.ZapConfig) (core zapcore.Core, options []zap.Option) {
	// 日志级别决策器
	levelEnabler := cfg.ZapLevelEnabler()
	consoleEncoder, filenCoder := cfg.ZapConsoleEncoder(), cfg.ZapFileEncoder()
	consoleSyncer, fileSyncer := cfg.ZapConsoleWriterSyncer(), cfg.ZapFileWriterSyncer()
	combineCore := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, levelEnabler),
		zapcore.NewCore(filenCoder, fileSyncer, levelEnabler),
	)
	return combineCore, cfg.BuildOptions(levelEnabler)
}

func initZapOutput(output []string) {
	for _, filePath := range output {
		fileutils.MustCreateDirAndFile(fileutils.JoinPath(filePath))
	}
}
