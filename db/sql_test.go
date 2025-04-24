package db

import (
	"testing"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


func TestTransactionDemo(t * testing.T) {
	TransactionDemo()
}
func TestPrepareInsertDemo(t *testing.T) {
	PrepareInsertDemo()
}

func TestPrepareQueryDemo(t *testing.T) {
	PrepareQueryDemo()
}
func TestDeleteRowDemo(t *testing.T) {
	DeleteRowDemo()
}

func TestUpdateRowDemo(t *testing.T) {
	UpdateRowDemo()
}

func TestInsertRowDemo(t *testing.T)  {
	InsertRowDemo()
}

func TestQueryMultiRowDemo(t *testing.T) {
	QueryMultiRowDemo()
}

func TestQueryRowDemo(t *testing.T) {
	QueryRowDemo()
}

func TestMysqlConn(t *testing.T) {
	err := initDB()

	if err != nil {
		log.Printf("init db failed, err:%v\n", err)
		return
	}
}

