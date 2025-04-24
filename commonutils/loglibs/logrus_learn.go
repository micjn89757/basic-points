package loglibs

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 创建一个新的logger实例。可以创建任意多个。
var log = logrus.New()

func defaultField() {
	// 将一些字段始终附加到应用程序
	// 返回一个Entry
	requestLogger := log.WithFields(logrus.Fields{
		"request_id": 111111,
		"user_ip":    1111111,
	})

	requestLogger.Info("demo")
}

func logLevel() {
	// 可以直接在logger上设置日志记录级别，然后只会记录具有该级别及以上级别任何内容的条目
	//log.SetLevel(logrus.InfoLevel)

	// 日志级别由低到高
	log.Trace("demo")
	log.Debug("demo")
	log.Info("demo")
	log.Warn("demo")
	log.Error("demo")

	// 记录完日志后会调用os.Exit()
	//log.Fatal("Bye.")

	// 记录完日志后会调用panic()
	//log.Panic("i'm bailing")

}

func outputDiffLoc() {
	// 设置日志输出为os.Stdout
	log.Out = os.Stdout

	// 设置日志输出为文件
	file, err := os.OpenFile("demo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	log.Out = file

	log.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("log info")
}

func baseUse() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("log info")
}
