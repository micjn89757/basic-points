package mylogger

import (
	"logger/mylogger/core"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestMylogger(t *testing.T) {
	file, err := os.OpenFile("./log/demo.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		t.Error(err)
	}

	EncodeLevel = core.InfoLevel
	logger := NewLogger(WithEncoding("json"), WithRotate("./log/demo.log", time.Hour * 24), WithWriter(os.Stderr, file))
	logger.Info("dddd", core.Uint64("ddd", 14))
	logger.Debug("ddd", core.Uint64("ddd", 16))
}


func BenchmarkLogger(b *testing.B) {
	file, _ := os.OpenFile("./log/demo.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)

	logger := NewLogger(WithEncoding("json"), WithRotate("./log/demo.log", time.Hour * 24), WithWriter(os.Stderr, file))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("dddd", core.Uint64("number", 16))
	}
}

// mylogger  10727 ns/op	544 B/op	7 allocs/op

func BenchmarkLogrus(b *testing.B) {
	file, _ := os.OpenFile("./log/demo.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	LogRus := logrus.New()
	LogRus.SetOutput(file)       //设置日志文件
	LogRus.SetReportCaller(true) //输出是从哪里调起的日志打印，这样日志里会多出func和file
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logrus.Infof("demo: %d", 12)
	}

}

//logrus	 21741 ns/op	 474 B/op	17 allocs/op