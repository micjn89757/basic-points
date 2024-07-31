package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	sync.Pool(并发安全)， 但是Pool.New func 可能会被并发调用
	程序员作为使用方不能对 Pool 里面的元素个数做假定
	我们不能对 Pool 池里的 cache 的元素个数做任何假设


	为什么 sync.Pool 不适合用于像 socket 长连接或数据库连接池?
	因为，我们不能对 sync.Pool 中保存的元素做任何假设，以下事情是都可以发生的：

	Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
	Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
	Pool 池里面的元素个数你无法知道；
	所以，只有的你的场景满足以上的假定，才能正确的使用 Pool 。sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担。划重点：临时对象。所以说，像 socket 这种带状态的，长期有效的资源是不适合 Pool 的
*/

var numCalcsCreated int32 

// 对象构造器
func createBuffer() any {
	// 必须注意要使用原子加，否则有并发问题
	atomic.AddInt32(&numCalcsCreated, 1)	// 构造一次+1
	buffer := make([]byte, 1024)
	return &buffer
}


func syncPool() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// 多goroutine并发测试
	numWorkders := 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkders)

	for i := 0; i < numWorkders; i++ {
		go func ()  {
			defer wg.Done()

			// 申请一个buffer实例
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)

			// 一定要释放实例
			defer bufferPool.Put(buffer)
		}()
	}

	wg.Wait()
	fmt.Println("buffer obj were created.\n", numCalcsCreated)
}