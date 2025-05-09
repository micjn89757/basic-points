package v1

import (
	"fmt"
	"sync"
)

// syncmap implement with map level

type SyncMap[T comparable, V any] struct {
	mp map[T]V
	mu sync.Mutex // or RWMutex
}

func NewSyncMap[K comparable, V any](cap uint) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		mp: make(map[K]V, cap),
	}
}

func (sm *SyncMap[T, V]) Set(key T, val V) (err error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("map set error: %s", p)
		}
	}()

	sm.mp[key] = val
	return
}

func (sm *SyncMap[T, V]) Get(key T) V {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	return sm.mp[key]
}
