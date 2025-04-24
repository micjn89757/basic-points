package main

import (
	"errors"
	"fmt"
)

// 可变参数通过切片实现的
func intSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum 
}

// 可变参数要放最后
func intSum2(x int, y ...int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}

	return sum 
}


// 多返回值用括号
func calcZ(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y 
	return 
}


// !定义函数类型
// 只要满足这个条件的函数都是calculation类型
type calculation func(int, int) int 

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// add  sub都是calculation类型变量
func funcType() int {
	var c calculation
	c = add 
	return c(1, 2)
}


// !高阶函数

// 作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 作为返回值
func do(s string) (func(int, int) int, error) {

	// 匿名函数
	add := func(x int, y int) int {
		return x + y
	}

	switch s {
	case "+":
		return add, nil 
	case "-":
		return sub , nil 
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}


// recover必须搭配defer使用
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}

	}()

	panic("panic in B")
}