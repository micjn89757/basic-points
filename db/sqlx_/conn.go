package sqlx_

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB 

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/atguigudb?charset=utf8&mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return nil
}


func init() {
	err := initDB()
	if err != nil {
		log.Printf("connect db failed, err:%v\n", err)
	}
	log.Printf("数据库连接成功..")
}