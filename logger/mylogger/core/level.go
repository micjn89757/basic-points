package core

import (
	"fmt"
	"log/slog"
)

// 存放log的级别
// 定义级别
type Level int8
const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel 
	ErrorLevel
	PanicLevel
	FatalLevel	// os.Exit(1)
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}


func (l Level) CapitalString() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}

	slog.Info()
}