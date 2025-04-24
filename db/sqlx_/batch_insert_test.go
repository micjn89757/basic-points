package sqlx_

import (
	"log"
	"testing"
	"time"
)


func TestBatchInsertEmps(t *testing.T) {
	emps := make([]*employee, 0, 2)
	emps = append(emps, &employee{
		Id: 1,
		Last_name: "sss",
		Email: "111111",
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})
	emps = append(emps, &employee{
		Id: 2,
		Last_name: "dddd",
		Email: "2222",
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})

	BatchInsertEmps(emps)
}

func TestBatchInsertIn(t *testing.T) {
	emps := make([]any, 0, 2)
	emps = append(emps, &employee{
		Id: 4,
		Last_name: "sss",
		Email: "11111",
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})
	emps = append(emps, &employee{
		Id: 5,
		Last_name: "iii",
		Email: "90980980",
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})

	BatchInsertIn(emps)
}

func TestBatchInserNamedExec(t *testing.T) {
	emps := make([]*employee, 0, 2)
	emps = append(emps, &employee{
		Id: 14,
		Last_name:  "ooo",
		Email: "12323w33", 
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})
	emps = append(emps, &employee{
		Id: 16,
		Last_name: "ppp",
		Email: "99888020",
		Hire_date: time.Now(),
		Job_id: "MK_REP",
	})

	err := BatchInsertNamedExec(emps)

	if err != nil {
		log.Printf("insert error, err:%v\n", err)
	}
}