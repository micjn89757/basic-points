package tip

import (
	"errors"
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

type MyConcurrentMap struct {
	mu sync.RWMutex
	mmap  map[int]int
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mmap: make(map[int]int),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mmap[k] = v
}


func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for {
		if _, ok := m.mmap[k]; ok {
			return m.mmap[k], nil
		}else {
			select {
			case <- time.After(maxWaitingDuration):
				return 0, errors.New("get time out")
			default:
			}
		}
	}
}