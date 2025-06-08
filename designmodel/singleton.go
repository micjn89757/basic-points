package designmodel

import (
	"sync"
)

type singleton struct {
	name string
}

var once sync.Once

var instance *singleton

func GetInstance() {
	once.Do(func() {
		instance = &singleton{
			name: "demo",
		}
	})
}

// 使用OnceFunc初始化，不需要初始化once
func GetInstanceFunc() {
	f := sync.OnceFunc(func() {
		instance = &singleton{
			name: "dddee",
		}
	})

	f()
}
