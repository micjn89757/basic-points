package db

import (
	"fmt"
	"log"
	"time"
)

type employee struct {
	id           int
	last_name    string
	email        string
	phone_number string
	hire_date    time.Time
}


// PrepareInsertDemo 预处理插入
func PrepareInsertDemo() {
	sqlStr := "insert into employees(last_name, email, hire_date, job_id) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Printf("Prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()
	ret, err := stmt.Exec("王五", "1980846456@qq.com", time.Now(), "AD_VP")


	if err != nil{
		log.Printf("insert failed")
		return
	}

	insertID, err := ret.LastInsertId()
	
	if err != nil {
		log.Printf("get insert id failed, err:%v\n", err)
		return
	}

	log.Printf("insert id : %v\n", insertID)
}

// PrepareQueryDemo 预处理查询
func PrepareQueryDemo() {
	sqlStr := "select last_name, email, hire_date from employees where employee_id >  ?"

	stmt, err := db.Prepare(sqlStr) // SQL语言预处理
	if err != nil {
		log.Printf("Prepare failed, err: %v\n", err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(100) // 发送数据部分给服务器端

	if err != nil {
		log.Printf("query Failed, err:%v\n", err)
		return 
	}
	
	defer rows.Close()

	//  循环读取结果集中的数据
	for rows.Next() {
		var emp employee
		err := rows.Scan(&emp.last_name, &emp.email, &emp.hire_date)

		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return
		}

		log.Printf("%v\n", emp)
	}
}

//QueryRowDemo 单行查询
func QueryRowDemo() {
	sqlStr := "select employee_id, last_name, email, phone_number, hire_date from employees where employee_id = ?"

	emp := &employee{}

	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 100).Scan(&emp.id, &emp.last_name, &emp.email, &emp.phone_number, &emp.hire_date)

	if err != nil {
		log.Printf("scan failed, err:%v\n", err)
		return
	}

	log.Printf("%#v, %T\n", emp, emp)

}


// QueryMultiRowDemo 多行查询
func QueryMultiRowDemo() {
	sqlStr := "select employee_id, last_name, hire_date from employees where employee_id > ?"

	rows, err := db.Query(sqlStr, 0)

	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		return
	}

	// 关闭rows释放持有的数据库连接
	defer rows.Close()

	// 循环读取结果中的数据
	for rows.Next() {
		emp := &employee{}

		err := rows.Scan(&emp.id, &emp.last_name, &emp.hire_date)

		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return
		}

		log.Printf("%#v\n", emp)
	}
}



// InsertRowDemo 插入数据
func InsertRowDemo() {
	sqlStr := "insert into employees(last_name, email, hire_date, job_id) values (?, ?, ?, ?)"

	res, err := db.Exec(sqlStr, "王五", "1980846456@qq.com", "2024-01-22", "AD_PRES")


	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return
	}


	insertID, err := res.LastInsertId() // 返回新插入的数据id
	if err != nil {
		log.Printf("get lastinsertid failed, err:%v\n", err)
		return
	}

	log.Printf("insert sucess, the id is %v\n",insertID)
}

// UpdateRowDemo 更新数据
func UpdateRowDemo() {
	sqlStr := "update employees set email = ? where employee_id = ?"
	ret, err := db.Exec(sqlStr, "34658789@163.com", 0)

	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return
	}
	
	n, err := ret.RowsAffected() // 操作影响的行数
	
	if err!= nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}

	log.Printf("update success, affected rows:%d\n", n)
}


// DeleteRowDemo 删除数据
func DeleteRowDemo() {
	sqlStr := "delete from employees where employee_id = ?"

	res, err := db.Exec(sqlStr, 0)

	if err != nil {
		log.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := res.RowsAffected() // 操作影响的行数

	if err != nil {
		fmt.Printf("get RowsAffected  failed, err: %v\n", err)
		return
	}

	log.Printf("delete success, affected rows:%d\n", n)
}