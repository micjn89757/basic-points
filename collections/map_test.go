package collections

import (
	"testing"
	// "runtime"
)

// TestMapInit map初始化方式
func TestMapInit(t *testing.T) {
	// myMap := make(map[int]int, 2)
	// myMap := make(map[int]int)  不指定容量，默认为0
	myMap := map[int]int{
		1:2,
		2:3,
	} //初始化并复制
	t.Errorf("%+v\n", myMap)
	t.Error("TEST")
}

// TestTraversalMap 遍历map
func TestTraversalMap(t *testing.T) {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
	}

	// 只打印key不打印value
	for k := range myMap {
		t.Logf("%+v", k)
	}
}

// 练习题，观察打印结果
func TestMapProblem(t *testing.T) {
	type Map map[string][]int 
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	t.Errorf("%+v, %p\n", s, s)
	m["djn"] = s 
	t.Errorf("%#v, %p\n", m["djn"], m["djn"]) 
	s = append(s[:1], s[2:]...)
	t.Errorf("%+v, %p\n", s, s)
	t.Errorf("%+v, %p\n", m["djn"], m["djn"])
	t.Error("TEST")
}


