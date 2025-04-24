package loglibs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var ginLog = logrus.New()

func init() {
	// 使用JSON格式打印日志而不是默认的
	ginLog.Formatter = &logrus.JSONFormatter{}

	file, err := os.OpenFile("demo.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ginLog.Info(err)
	}

	ginLog.Out = file

	gin.SetMode(gin.ReleaseMode) // 程序为release模式
	gin.DefaultWriter = ginLog.Out

	// 设置打印级别
	ginLog.Level = logrus.InfoLevel
}

func logSever() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		ginLog.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Warn("A group of walrus emerges from the ocean")

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run(":8080")
}
