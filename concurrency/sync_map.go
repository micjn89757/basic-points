package concurrency

import (
	"fmt"
	"sync"
)

// 普通的map并发不安全
var smp = sync.Map{}
// 使用sync.Map也会存在问题(脏写)
func mapInc(mp *sync.Map, key int) { // 注意必须传sync.Map的指针，要修改结构体，必须传指针
	if oldValue, exists := mp.Load(key); exists {
		mp.Store(key, oldValue.(int)+1)
	} else {
		mp.Store(key, 1)
	}

}

func syncMap() {
	const P = 1000
	const key = 8
	wg := sync.WaitGroup{}
	wg.Add(P)

	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done() 
			mapInc(&smp, key)
		}()
	}
	wg.Wait()
	value, _ := smp.Load(key)
	fmt.Println(value)
}