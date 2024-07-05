package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB  // 内部维护一个连接池

var conf Config

var logger *zap.Logger

// toml文件配置
type mysql struct {
	Host string `toml:"host"`
	Username string	`toml:"username"`
	Password string	`toml:"password"`
	ConnParams param `toml:"connparams"`
}

type param struct {
	Charset string 
	ParseTime bool
}

type Config struct {
	MSQL *mysql `toml:"mysql"`
}


func init() {
	initLogger()
	connectAll()
}

func connectAll() error {
	var err error
	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Info("数据库连接中...")
	_, err = toml.DecodeFile("config.toml", &conf)

	if err != nil {
		sugar.Infof("err:%v", zap.Error(err))
		return err 
	}


	dbname := "atguigudb"
	msqldsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", 
						conf.MSQL.Username, 
						conf.MSQL.Password, 
						conf.MSQL.Host,
						dbname,
						conf.MSQL.ConnParams.Charset)	
	
	if conf.MSQL.ConnParams.ParseTime {
		msqldsn += "&parseTime=True"
	}

	db, err = sqlx.Open("mysql", msqldsn) // 创建数据库句柄，不会校验账号密码是否正确

	if err != nil {
		sugar.Infof("参数格式错误: %v", err)
		return err
	}

	// defer db.Close()  // 一定要写在错误处理之后

	err = db.Ping() // 尝试与数据库建立连接，校验dsn
	if err != nil {
		sugar.Infof("连接失败:%v", err)
		return err 
	}

	sugar.Infof("数据库连接成功")
	return nil
}


// initLogger 初始化日志
func initLogger() {
	var err error
	logger, err = zap.NewProduction()

	if err != nil {
		panic("日志初始化失败")
	}
}


