package json

import "time"

type User struct {
	Name     string
	Age      int
	height   float32 	// 不可导出成员不会被序列化
	BirthDay time.Time	// 格式：2024-07-24T20:14:11.7074482+08:00
	CreateAt MyDate		// 格式: 2024-07-24
}

// 标准库json序列化背后使用的核心技术是反射。通过反射可以在运行时动态获得结构体成员变量的名称、json(tag)、是否可导出，可以获取成员变量的值，还可以调用结构体的方法（比如MarshallJSON和UnmarshallJSON）
