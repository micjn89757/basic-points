/*
日志切分和滚动相关
*/
package mylogger

import (
	"fmt"
	"logger/mylogger/core"
	"os"
	"sync"
	"time"
)

// TODO: 日志滚动功能 maxAge rotationCount
// 支持日志切割的logger
type RotateLogger struct {
	initTime 		time.Time
	FileName 		string 
	RotationTime	time.Duration	// 隔多久分割一次
	MaxAge			time.Duration	// TODO:设置文件清理前的最长保存时间
	RotationCount	uint64			// TODO: 设置文件清理前最多保存的个数

	sync.Mutex
}


func NewRotateLog(fileName string, rotationTime time.Duration) *RotateLogger {
	rl := &RotateLogger{
		initTime: time.Now(),
		FileName: fileName,
		RotationTime: rotationTime,
	}

	return rl
}


// CheckAndCangeLogFile 检查是否需要对日志进行分割，需要则分割，支持并发调用
func(rl *RotateLogger) CheckAndChangeLogFile(ws core.WriteSyncer) core.WriteSyncer {
	rl.Lock()
	defer rl.Unlock()
	now := time.Now()

	if duration := now.Sub(rl.initTime); duration <= rl.RotationTime {
		return  nil // 不需要分割
	}

	// 关闭老的日志输出
	ws.Close()

	// 给老的日志文件加上日期后缀
	postFix := now.Add(-24 * time.Hour).Format("20060102")	// 昨天的日期

	// 重命名文件
	if err := os.Rename(rl.FileName, rl.FileName + "." + postFix); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("append date postfix %s to log file %s failed: %v\n", postFix, rl.FileName, err)) // 如果Logger本身出错，则把错误信息打到标准错误输出里

		return nil
	}

	// 打开新的文件日志
	if logOut, err := os.OpenFile(rl.FileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0664); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("create log file %s failed: %v\n",rl.FileName, err))
		return nil
	} else {
		rl.initTime = time.Now()
		return logOut
	}
}