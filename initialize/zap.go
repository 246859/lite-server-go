package initialize

import (
	"strings"

	"go.uber.org/zap"
	"liteserver/config"
)

// InitZap
// @Date: 2023-01-09 14:52:27
// @Description: Zap初始化
// @Param: config *config.ZapConfig
// @Return: *zap.Logger
func InitZap(config *config.ZapConfig) *zap.Logger {
	zapConfig := config.GetConfig()
	logger := zap.Must(zapConfig.Build())
	logger.Info("应用日志初始化成功...",
		zap.String("日志级别", zapConfig.Level.String()),
		zap.String("输出目录", strings.Join(zapConfig.OutputPaths, ",")))
	return logger
}
