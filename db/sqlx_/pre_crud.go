package sqlx_

import (
	"log"
	"time"
)

func prepareQuery() {
	sqlStr := "select employee_id, last_name, email, hire_date, job_id from employees where employee_id > ?"

	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Printf("prepare failed, err: %v\n", err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(100)

	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		return
	}

	// 循环读取结果集中的数据
	// 调用rows.Next如果返回false并且没有更多结果集，rows会自动close
	for rows.Next() {
		var emp employee
		err := rows.Scan(&emp.Id, &emp.Last_name, &emp.Email, &emp.Hire_date, &emp.Job_id)

		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return
		}

		log.Printf("%v\n", emp)

	}
}

func prepareInsert() {
	sqlStr := "insert into employees(employee_id, last_Name, email, hire_date, job_id) values(?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Exec(10, "ddd", "1089775@qq.com", time.Now(), "MK_REP")

	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}

	insertId, err := res.LastInsertId()

	if err != nil {
		log.Printf("get insert id failed, err:%v\n", err)
		return
	}

	log.Printf("insert success, insert id: %v\n", insertId)
}