package test

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZap(t *testing.T) {
	if zap.L().Level() == zapcore.InvalidLevel {
		fmt.Println("nopcore")
	}
}
