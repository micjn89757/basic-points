package tip

import (
	"context"
	"sync"
	"time"
)

/*
实现一个map，
面向高并发
只存在插入和查询操作O(1)
查询时，若key存在，直接返回val；若key不存在，阻塞直到key val对被放入后， 获取val返回
等待指定时长仍未被放入，返回超时错误
写出真实代码，不能有死锁或者panic风险
*/

// 防止channel重复关闭，方案三
type MyChan struct {
	sync.Once
	ch chan struct{}
}

func NewMyChan() *MyChan {
	return &MyChan{
		ch: make(chan struct{}),
	}
}

func (mc *MyChan) Close() {
	mc.Do(func() {
		close(mc.ch)
	})
}

// 读写锁可能会有问题
type MyConcurrentMap struct {
	mu sync.Mutex
	mmap  map[int]int
	keyToCh map[int]chan struct{}
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mmap: make(map[int]int),
		keyToCh: make(map[int]chan struct{}),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mmap[k] = v

	// 读goroutine负责传递信号放入keyToCh中，唤醒阻塞的读协程，要注意的是一个struct{}{}只能唤醒读阻塞队列中的一个读协程，要唤醒全部，所以要通过关闭channel，关闭channel会唤醒所有的阻塞goroutine
	if m.keyToCh[k] == nil {
		return 
	} 

	// 防止重复关闭
	// 方案一，有值的情况下表示channel已经被关闭了，陷入阻塞的情况下，表示没有关闭过
	select {
	case <- m.keyToCh[k]:
		return 
	default:
		close(m.keyToCh[k])
	}

	// 方案二，直接删除
	// close(m.keyToCh[k])
	// delete(m.keyToCh, k)

	// 方案三，设置一个不可重复关闭的channel，单例模式


	// 关闭channel，并没清除掉m.keyTocCh[k],如果此时第二个写goroutine到达，还会继续关闭
}


func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 有值就返回
	v, ok := m.mmap[k]
	if ok {
		return v, nil
	}
	
	// 这里的问题是，并发场景多个goroutine读，可能会创建多次，所以要检查，这里不用读写锁的原因在于，这里读的时候其实也写了，读读共享可能导致多个goroutine同时创建
	ch, ok := m.keyToCh[k]
	if !ok {
		m.keyToCh[k] = make(chan struct{})  
	} 

	// 设置挂起超时
	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()

	m.mu.Unlock()  // 挂起之前一定要解锁，否则会锁住整个map，不能读也不能写

	select {
	case <- ch: // 等待唤醒
	case <- tCtx.Done():
		return -1, tCtx.Err()
	}

	// 读之前还要加锁
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.mmap[k], nil
}