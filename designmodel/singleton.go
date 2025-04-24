package designmodel

import (
	"sync"
)


type singleton struct{}

var once sync.Once

var instance *singleton

func GetInstance() *singleton {
	once.Do(func ()  {
		instance = &singleton{}
	})

	return instance
}