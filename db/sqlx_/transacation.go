package sqlx_

import (
	"errors"
	"log"
)

// sqlx事务操作
func transactionDemo() (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		log.Printf("begin trans failed, err:%v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after rollback
		} else if err != nil {
			log.Printf("rollback\n")
			tx.Rollback()
		} else {
			err = tx.Commit() // err is nil; if commit returns error update err
			log.Printf("commit")
		}
	}()

	sqlStr1 := `update employees set last_name="ppp" where employee_id = ?`

	res, err := tx.Exec(sqlStr1, 0)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := `update employees set last_name="www" where employee_id = ?`
	res, err = tx.Exec(sqlStr2, 1)
	if err != nil {
		return err
	}

	n, err = res.RowsAffected()
	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}

	return err
}
