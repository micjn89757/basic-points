package commonutils

// 标准日志库， 实践中常用一些第三方日志库比如loguru,zap等
import (
	"log"
	"os"
)

// 配置logger
func LoggerSet() {
	// log标准库提供了一些flag选项(常量)
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// // log.Println("这是一条普通日志")
	// // log.Println(log.Flags()) //会返回标准logger的输出配置, 以flags bit的和返回

	// // 配置日志前缀
	// log.SetPrefix("[djn]")
	// // 查看日志输出前缀
	// // log.Prefix() 返回string
	log.Println("日志信息")

	// 配置日志输出位置
	// 默认是标准错误输出到控制台
	// 输出到文件
	// logFile, err := os.OpenFile("a.log", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0666)
	// if err != nil {
	// 	// 处理错误
	// 	return
	// }

	// // 设置log输出到文件
	// log.SetOutput(logFile)

	// log.Println("输出到文件")
	
}

// 自定义logger对象
func LoggerCreate() {
	// 第一个参数是输出目的，第二个参数是前缀，第三个是日志属性
	logger := log.New(os.Stdout, "<New>", log.Lshortfile | log.Ldate | log.Ltime)
	logger.Println("自定义日志对象")
}

// 使用标准logger通常把配置操作写到init中，在被导入的时候自动执行配置, 对整个包的log都生效
func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("[djn]")
	
	logFile, err := os.OpenFile("a.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Println("open log file error", err)
		return
	}

	log.SetOutput(logFile)

	log.Println("日志打印")
}