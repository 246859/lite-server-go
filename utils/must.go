package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// MustOrPanic
// @Date 2023-01-12 18:23:11
// @Param f func() error
// @Description: 必须完成某个操作，否则就panic
func MustOrPanic(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

// MustOrLogPanic
// @Date 2023-01-12 19:15:33
// @Param f func() error
// @Param msg string
// @Param field ...zap.Field
// @Description: 必须完成某一个操作，否则就panic并记录日志
func MustOrLogPanic(f func() error, msg string, field ...zap.Field) {
	if err := f(); err != nil {
		logger := zap.L()
		// 判断是否是nopcore
		if logger.Level() == zapcore.InvalidLevel {
			log.Panicln(msg, err)
		} else {
			logger.Panic(msg, field...)
		}
	}
}
