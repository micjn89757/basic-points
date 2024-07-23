package main

import (
	"errors"
	"time"
)

// employee schema
type employee struct {
	EmployeeId int `db:"employee_id"`
	FirstName string `db:"first_name"`
	LastName string	`db:"last_name"`
	Email string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	HireDate time.Time	`db:"hire_date"`
	JobID string `db:"job_id"`
	Salary float64 `db:"salary"`
	CommissionPCT float64 `db:"commission_pct"`
	ManagerID int  `db:"manager_id"`
	DepartmentID int `db:"department_id"`
}

// queryRowDemo 查询单条数据示例
func queryRowDemo() {
	sugar := logger.Sugar()
	defer sugar.Sync()
	var err error

	sqlStr := "select employee_id, first_name, email from employees where id = ?"
	var emp employee

	err = db.Get(&emp, sqlStr, 100)

	if err != nil {
		sugar.Infof("err: %v", err)
		return 
	}

	sugar.Info(emp)
}

// queryMultiRowDemo 查询多条数据示例
func queryMultiRowDemo() {
	sugar := logger.Sugar()
	defer sugar.Sync()
	var err error 

	sqlStr := "select employee_id, first_name, last_name from employees"

	var emps []employee
	err = db.Select(&emps, sqlStr)

	if err != nil {
		sugar.Infof("err: %v", err)
		return 
	}

	sugar.Info(emps)
}

// insertRowDemo 插入数据
func insertRowDemo() {
	sugar := logger.Sugar()
	defer sugar.Sync()
	sqlStr := "insert into employees(employee_id, last_name, email, hire_date, job_id) values(?, ?, ?, ?, ?)"
	ret, err := db.Exec(sqlStr, 209, "djn", "198@qq.com", time.Now(), "AC_ACCOUNT")
	if err != nil {
		sugar.Errorf("insert err: %w", err)
		return 
	}

	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		sugar.Errorf("get lastinsert ID failed, err:%v", err)
		return 
	}
	sugar.Infof("insert success, the id is %d", theID)
}


// updateRowDemo 更新数据
func updateRowDemo() {
	sugar := logger.Sugar()
	defer sugar.Sync()
	sqlStr := "update employees set last_name = ? where employee_id = ?"
	ret, err := db.Exec(sqlStr, "ddd", 209)
	
	if err != nil {
		sugar.Errorf("update failed, %w", err)
		return 
	}

	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		sugar.Errorf("get rowsaffected failed, err:%w", err)
	}

	sugar.Infof("update success, affected rows:%d", n)
}

// deleteRowDemo 删除数据
func deleteRowDemo() {
	sugar := logger.Sugar()
	defer sugar.Sync()

	sqlStr := "delete from employees where employee_id = ?"
	ret, err := db.Exec(sqlStr, 209)

	if err != nil {
		sugar.Errorf("delete failed, err:%w", err)
		return 
	}

	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		sugar.Errorf("get rowsaffected failed, err:%w", err)
		return 
	}

	sugar.Infof("delete success, affected rows:%d", n)
}

// transactionDemo2() 
func transactionDemo2() (err error) {
	sugar := logger.Sugar()
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		sugar.Infof("begin trans failed, err:%w", err)
		return err 
	}

	// 什么时候回滚和提交事务
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			sugar.Info("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			sugar.Info("commit")
		}
	}()

	sqlStr1 := "insert into employees(employee_id, last_name, email, hire_date, job_id) values(?, ?, ?, ?, ?)"
	
	ret, err := tx.Exec(sqlStr1, 209, "djn", "198@11.com", time.Now(), "AC_ACCOUNT")
	if err != nil {
		return err 
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "update employees set last_name = ? where employee_id = ?"
	ret, err = tx.Exec(sqlStr2, "ddd", 209)
	if err != nil {
		return err
	}

	n, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}

	return err
}