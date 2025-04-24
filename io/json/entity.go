package json

import (
	"time"
)


type User struct {
	Name string 
	Age int `json:"age"`
	height float32
	Birthday time.Time // 默认的time.Time格式
	CreatedAt MyDate //格式: 2023-09-29 ——需要自己实现
}