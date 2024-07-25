package core

import (
	"runtime"
	"time"
)

// 存放log信息的实体
type Entry struct {
	Level	Level		// 第一个字段是日志级别
	Time	string	// 第二个字段是时间
	Message	string 		// 第三个字段是打印的消息
	Caller	*CallerEntry // 调用的函数名文件名以及所在行
}


// 存储调用堆栈信息
type CallerEntry struct {
	File	string 
	FC 		string		// function name
	Line	int
}

func NewEntry(level Level, time time.Time, msg string) *Entry {
	funcName, file, line, ok := runtime.Caller(2)
	if ok {
		ce := &CallerEntry{
			File: file,
			FC: runtime.FuncForPC(funcName).Name(),
			Line: line,
		}
		
		return &Entry{
			Level: level,
			Time: time.Format("2006-01-02 15:04:05"),
			Message: msg,
			Caller: ce,
		}
	} else {
		return nil 
	}
}

// func getFileAndLineNo() {
// 	runtime.Caller(3) // 返回第三层调用堆栈
// }