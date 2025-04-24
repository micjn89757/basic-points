package sqlx_

import (
	"log"
	"testing"
)

func TestQueryAndOrderByIDs(t *testing.T) {
	emps, err :=QueryAndOrderByIDs([]int{14, 11, 10})

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	log.Printf("%v\n", emps)
}

func TestQueryByIDs(t *testing.T) {
	emps, err :=QueryByIDs([]int{11, 13, 10})

	if err != nil {
		return
	}

	log.Printf("%v\n", emps)
}

func TestNamedQuery(t *testing.T) {
	namedQuery()
}

func TestNamedExec(t *testing.T) {
	namedExec()
}

func TestQueryRowDemo(t *testing.T) {
	queryRowDemo()
}


func TestQueryMultiRowDemo(t *testing.T) {
	queryMultiRowDemo()
}

func TestInsertRowDemo(t *testing.T) {
	insertRowDemo()
}

func TestUpdateRowDemo(t *testing.T) {
	updateRowDemo()
}

func TestDeleteRowDemo(t *testing.T) {
	deleteRowDemo()
}