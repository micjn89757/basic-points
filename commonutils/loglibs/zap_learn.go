package loglibs

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// var logger *zap.logger

// InitLogger 选择不同方式初始化logger
func InitLogger() {
	// 不同初始化方式的默认自带的信息不同
	loggerE := zap.NewExample() // 自带level字段，输出格式是json
	loggerE.Info("logggerE", zap.String("init", "newExample"))

	loggerD, err := zap.NewDevelopment() // 自带日期
	if err != nil {
		fmt.Print(err)
	}
	loggerD.Info("loggerD")
	loggerD.Debug("loggerD")
}

// LoggerChoose 
func LoggerChoose() {
	// 两种logger默认都是打印json格式
	// 在不重视内存分配的上下文中可以使用sugared logger
	sugar := zap.NewExample().Sugar()

	defer sugar.Sync() // 刷新缓冲区

	sugar.Infow("failed to fetch URL", "url", "https://example.com", "attempt", 3, "backoff", time.Second)

	sugar.Infof("failed to fetch URL: %s", "http://example.com")

	// 在重视内存分配的上下文中使用Logger, 比SugarLogger更快
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("failed to fetch URL", zap.String("url", "http://example.com"), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
}

