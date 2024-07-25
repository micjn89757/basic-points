package gorm_

import (
	"database/sql"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	sqlDB *sql.DB
	sqlDBOnce sync.Once
)


func init() {
	initDBOnce()
}


func initDBOnce() {
	sqlDBOnce.Do(func() {
		dsn := "root:123456@tcp(192.168.197.133:3306)/atguigudb?charset=utf8mb4&parseTime=True&loc=Local"
		atguiguDB, err := gorm.Open(mysql.New(
			mysql.Config{
				DSN: dsn,
				DefaultStringSize: 256, // string 类型字段默认长度
				// DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			}), &gorm.Config{
				PrepareStmt: true, // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
				// NamingStrategy: schema.NamingStrategy{
				// 	SingularTable: true,// 使用单数表名
				// },
			})
		if err != nil {
			panic(err)
		}

		// 维护连接池
		sqlDB, err = atguiguDB.DB()

		if err != nil {
			panic(err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
}

func getDB() *sql.DB {
	return sqlDB
}