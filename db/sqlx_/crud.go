package sqlx_

import (
	"database/sql/driver"
	"log"
	"time"
	"fmt"
	"strings"
	"github.com/jmoiron/sqlx"
)

type employee struct {
	Id        int       `db:"employee_id"`
	Last_name string    `db:"last_name"`
	Email     string    `db:"email"`
	Hire_date time.Time `db:"hire_date"`
	Job_id    string    `db:"job_id"`
}

func (emp *employee) Value() (driver.Value, error) {
	return []any{emp.Id, emp.Last_name, emp.Email, emp.Hire_date, emp.Job_id}, nil
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int)(emps []employee, err error){
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id)) // 转换成字符串并append
	}

	query, args, err := sqlx.In("SELECT employee_id, last_name, email, hire_date job_id FROM employees WHERE employee_id IN (?) ORDER BY FIND_IN_SET(employee_id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	log.Printf("query: %v\n", query)
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	log.Printf("query: %v\n", query)
	err = db.Select(&emps, query, args...)
	return
}


// 实现In查询
func QueryByIDs(ids []int) (emps []employee, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT employee_id, last_name, email, hire_date, job_id FROM employees WHERE employee_id IN (?)", ids)

	if err != nil {
		return
	}

	log.Printf("query: %v\n", query)
	log.Printf("args: %v\n", args)
	// sqlx.In 返回带 `?`bindvar的查询语句, 使用Rebind()重新绑定
	query = db.Rebind(query)
	err = db.Select(&emps, query, args...)
	return
}

// NamedQuery
func namedQuery() {
	sqlStr := "select last_name, email, hire_date, job_id from employees where employee_id =:id "

	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]any{
		"id": 100,
	})

	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var emp employee
		err := rows.StructScan(&emp)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			continue
		}

		log.Printf("emp:%v\n", emp)
	}

	emp := employee{
		Id: 100,
	}

	// 使用结构体命名查询，根据结构体字段的db tag进行映射
	rows, err = db.NamedQuery(sqlStr, emp)

	if err != nil {
		log.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var emp employee
		err := rows.StructScan(&emp)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			continue
		}

		log.Printf("emp:%v\n", emp)

	}
}

// NameExec 绑定SQL语句与结构体或map中的同名字段
func namedExec() {
	sqlStr := "insert into employees(last_name, email, hire_date, job_id) values(:last_name, :email, :hire_date, :job_id)"
	ret, err := db.NamedExec(sqlStr, map[string]any{
		"last_name": "qimi",
		"email":     "18888",
		"hire_date": time.Now(),
		"job_id":    "AD_VP",
	})

	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		log.Printf("get rowsaffected failed,err:%v\n", err)
		return
	}

	log.Printf("changed rows:%v\n", n)
}

// insertRowDemo 插入数据
func insertRowDemo() {
	sqlStr := "insert into employees(employee_id, last_name, email, hire_date, job_id) values (?, ?, ?, ?, ?)"

	ret, err := db.Exec(sqlStr, 1, "jjj", "198084656@qq.com", time.Now(), "AD_VP")

	if err != nil {
		log.Printf("insert error, err:%v\n", err)
		return
	}

	id, err := ret.LastInsertId()

	if err != nil {
		log.Printf("get insert id err: %v\n", err)
		return
	}

	log.Printf("insert success, thed id is %d\n", id)
}

// updateRowDemo 更新数据
func updateRowDemo() {
	sqlStr := `update employees set last_name = "ggg" where employee_id = ?`
	ret, err := db.Exec(sqlStr, 0)

	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()

	if err != nil {
		log.Printf("get rows affected failed, err:%v\n", err)
		return
	}

	log.Printf("update success, affected rows:%d\n", n)
}

// deleteRowDemo 删除数据
func deleteRowDemo() {
	sqlStr := "delete from employees where employee_id = ?"
	ret, err := db.Exec(sqlStr, 0)
	if err != nil {
		log.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		log.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}

	log.Printf("delete success, affected rows:%v\n", n)
}

// queryMultiRowDemo 查询多行数据
func queryMultiRowDemo() {
	defer func() {
		err := recover()

		if err != nil {
			log.Printf("发生panic, %s\n", err)
		}
	}()

	sqlStr := "select last_name, email, hire_date, job_id from employees where employee_id > ?"
	emps := make([]employee, 0)

	err := db.Select(&emps, sqlStr, 0)
	if err != nil {
		log.Printf("query failed, err: %v\n", err)
		return
	}

	log.Printf("emps:%#v\n", emps)
}

// queryRowDemo 查询单行数据
func queryRowDemo() {
	defer func() {
		err := recover()

		if err != nil {
			log.Printf("发生panic, %s\n", err)
		}

	}()

	sqlStr := "select last_name, email, hire_date, job_id from employees where employee_id = ?"
	// var emp employee
	emp := &employee{}

	err := db.Get(emp, sqlStr, 100)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		return
	}

	log.Printf("%#v\n", emp)
}
