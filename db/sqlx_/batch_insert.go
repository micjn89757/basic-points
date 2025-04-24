package sqlx_

import (
	"fmt"
	"strings"
	"github.com/jmoiron/sqlx"
	"log"
)

// BatchInsertEmps 自行构造批量插入的语句
func BatchInsertEmps(emps []*employee) error {
	// 存放(?, ?, ?, ?, ?)的slice
	valueStrings := make([]string, 0, len(emps))
	// 存放values的slice
	valueArgs := make([]any, 0, 5*len(emps))

	// 遍历emps准备相关数据
	for _, u := range emps {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, u.Id)
		valueArgs = append(valueArgs, u.Last_name)
		valueArgs = append(valueArgs, u.Email)
		valueArgs = append(valueArgs, u.Hire_date)
		valueArgs = append(valueArgs, u.Job_id)
	}

	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO employees (employee_id, last_name, email, hire_date, job_id) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertIn 使用sqlx.In 实现批量插入，需要结构体实现driver.Value接口
func BatchInsertIn(emps []any) error { // 注意In的第二个参数是any类型（interface{}）
	query, args, _ := sqlx.In("INSERT INTO employees (employee_id, last_name, email, hire_date, job_id) VALUES (?), (?)",
		emps..., // 如果arg实现了driver.Valuer, sqlx.In会通过调用Value展开它
	)

	log.Printf("sql string: %v\n", query)
	log.Printf("args: %v\n", args) // 生成的args(替换？的数据)

	_, err := db.Exec(query, args...)

	if err != nil {
		log.Printf("insert falied, err:%v\n", err)
		return err
	}

	return err
}

// BatchInsertNamedExec 使用NamedExec实现批量插入
func BatchInsertNamedExec(emps []*employee) error {
	_, err := db.NamedExec("INSERT INTO employees (employee_id, last_name, email, hire_date, job_id) VALUES (:employee_id, :last_name, :email, :hire_date, :job_id);", emps)

	return err
}