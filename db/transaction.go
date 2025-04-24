package db

import "log"

// TransactionDemo 事务示例
func TransactionDemo() {
	tx, err := db.Begin() // 开启事务

	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		log.Printf("begin failed, err:%v\n", err)
		return
	}

	sqlStr1 := `update employees set last_name="diu" where employee_id = ? `
	ret1, err := tx.Exec(sqlStr1, 0)
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec sql1 failed, err:%v\n", err)

		return
	}

	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		log.Printf("exec ret1.RowsAffeted() failed, err:%v\n", err)
		return
	}

	sqlStr2 := `Update employees set last_name = "jjj" where employee_id=?`
	ret2, err := tx.Exec(sqlStr2, 0)

	if err != nil {
		tx.Rollback()
		log.Printf("exec sql2 failed, err:%v\n", err)
		return
	}

	affRow2, err := ret2.RowsAffected()

	if err != nil {
		tx.Rollback()
		log.Printf("exec ret2.RowsAffected failed, err:%v\n", err)
		return 
	}

	log.Printf("af1:%v\n, af2:%v\n", affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		log.Print("事务提交了...")
		tx.Commit() //提交事务
	}else {
		tx.Rollback()
		log.Printf("事务回滚了...")
	}

}