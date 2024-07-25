package model

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID 				int32				`gorm:"primarykey;column:employee_id"`
	FirstName		string 
	LastName		*sql.NullString
	Email			*sql.NullString 	`gorm:"unique"`
	PhoneNumber		*sql.NullString 
	HireDate		time.Time
	JobID			string				
	Salary			*sql.NullFloat64
	CommissionPct	*sql.NullFloat64
	ManagerID		*sql.NullInt32
	DepartmentID	*sql.NullInt32
}

// 规定表名，实现Tabler接口
func (Employee) TableName() string {
	return "employees"
}