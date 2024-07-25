/*
zap的使用
*/
package zap_

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func initZap() {
	Logger, _ = zap.NewProduction()
	// Logger.Info("", zap.Int16())
	zap.L()
}

func encoder() zapcore.EncoderConfig {
	return zap.NewDevelopmentEncoderConfig()
}