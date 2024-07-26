package mylogger

import (
	"io"
	"logger/mylogger/core"
	"os"
	"strconv"
	"strings"

	"time"
)

var (
	EncodeLevel core.Level = core.InfoLevel // 打印的日志级别，比如默认不会将debug级别打印到日志中
)


type Logger struct {
	encoding		string 				// 编码方式，例如json

	rotate 			bool				// 是否支持日志切割
	rotateLog 		*RotateLogger

	ws 				core.WriteSyncer	// 日志输出位置
}


type LoggerOption func(*Logger)

func WithEncoding(encoding string) LoggerOption {
	return func(l *Logger) {
		l.encoding = encoding
	}
}

func WithWriter(writers ...io.Writer) LoggerOption {
	writer := io.MultiWriter(writers...)
	return func(l *Logger) {
		l.ws = core.AddSync(writer)
	}
}

// 需要传入分割日志的文件路径
func WithRotate(fileName string, rotateTime time.Duration) LoggerOption {
	return func(l *Logger) {
		l.rotate = true
		l.rotateLog = NewRotateLog(fileName, rotateTime)
	}
}

func NewLogger(options ...LoggerOption) *Logger {
	logger := &Logger{
		encoding: "json",	// 默认json
		ws: os.Stderr,		// 默认输出到Stderr
	}
	
	for _, opt := range options {
		opt(logger)
	}

	return logger
}


func (logger *Logger) Info(msg string, field ...core.Field) {
	// 检查日志级别，不够则不打印
	if core.InfoLevel >= EncodeLevel {
		// 检查是否需要日志分割
		if logger.rotate {
			if logOut := logger.rotateLog.CheckAndChangeLogFile(logger.ws); logOut != nil {
				logger.ws = core.AddSync(io.MultiWriter(logOut, os.Stderr))
			}
		}

		// 创建打印的基本信息
		entry := core.NewEntry(core.InfoLevel, time.Now(), msg)

		// 日志输出
		logger.write(entry, field...)
	}
}

func (logger *Logger) Debug(msg string, field ...core.Field) {
	if core.DebugLevel>= EncodeLevel {
		// 检查是否需要日志分割
		if logger.rotate {
			if logOut := logger.rotateLog.CheckAndChangeLogFile(logger.ws); logOut != nil {
				logger.ws = core.AddSync(io.MultiWriter(logOut, os.Stderr))
			}
		}

		// 创建打印的基本信息
		entry := core.NewEntry(core.DebugLevel, time.Now(), msg)

		// 日志输出
		logger.write(entry, field...)
	}
}



// 写入日志，不使用marshallJSON，而是拼接字符串
func (logger *Logger) write(entry *core.Entry, fields ...core.Field) {
	var builder strings.Builder

	builder.WriteRune('{')

	// 打印level 
	level := entry.Level.CapitalString()
	builder.WriteRune('"')
	builder.WriteString("level")
	builder.WriteString("\":")
	builder.WriteString(level)
	builder.WriteRune(',')

	// 打印time
	tim := entry.Time
	builder.WriteRune('"')
	builder.WriteString("time")
	builder.WriteString("\":")
	builder.WriteRune('"')
	builder.WriteString(tim)
	builder.WriteRune('"')
	builder.WriteRune(',')


	// 打印msg
	msg := entry.Message
	builder.WriteRune('"')
	builder.WriteString("msg")
	builder.WriteString("\":")
	builder.WriteRune('"')
	builder.WriteString(msg)
	builder.WriteRune('"')
	builder.WriteRune(',')
	
	for _, field := range fields {
		// TODO： 根据不同的类型
		builder.WriteRune('"')
		builder.WriteString(field.Key)
		builder.WriteString("\":")
		builder.WriteString(field.String) // 数字不需要双引号
		builder.WriteRune(',')
	}


	// 打印Caller
	caller := entry.Caller
	builder.WriteRune('"')
	builder.WriteString("caller")
	builder.WriteString("\":")
	builder.WriteRune('"')
	builder.WriteString(caller.FC)
	builder.WriteRune(':')
	builder.WriteString(strconv.Itoa(caller.Line))
	builder.WriteRune('"')

	
	builder.WriteRune('}')
	// builder.WriteByte('\n')
	logger.ws.Write([]byte(builder.String()))
}

