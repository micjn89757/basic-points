/*
实现一个自旋锁，适用于使用者保持锁时间比较短，因此才选择自旋而不是睡眠
*/
package mylock

import (
	"runtime"
	"sync"
	"sync/atomic"
)


type SpinLock uint32

const maxBackoff = 16	// 自旋锁

func (sl *SpinLock) Lock() {	// 上锁
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) { // 这个操作会尝试将0->1，但是如果发现spinLock已经为1，就代表已经被其他协程占用了
		for i := 0; i < backoff; i++ {	// 开始自旋，调用者不会睡眠，而一直循环等待
			runtime.Gosched()	// goroutine让出调度，因为此时可能比较繁忙
		}

		// 增加自选次数
		if backoff < maxBackoff {
			backoff = backoff << 1  // 增加一倍
		}
	}
}


func (sl *SpinLock) Unlock() {	// 解锁
	atomic.StoreUint32((*uint32)(sl), 0)  // 持有者直接解锁即可
}


func NewSpinLock() sync.Locker {
	return new(SpinLock)
}

