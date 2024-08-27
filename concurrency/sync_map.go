package concurrency

import (
	mylock "concurrency/my-lock"
	"fmt"
	"sync"
	"unsafe"

	farmhash "github.com/leemcloughlin/gofarmhash"
)

// 普通的map并发不安全
// sync.Map是一个结构体

var smp = sync.Map{}
// var ml = sync.Mutex{}
// 使用sync.Map也会存在问题(脏写)
func mapInc(mp *sync.Map, key int) { // 注意必须传sync.Map的指针，要修改结构体，必须传指针
	// ml.Lock()
	// defer ml.Unlock()
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

	// 使用ConcurrentMap
	// var cmp  = NewConcurrentHashMap[int64](8, 1000)
	// var ml sync.Mutex
	// for i := 0; i < P; i++ {
	// 	go func() {
	// 		defer wg.Done()

	// 		ml.Lock()
	// 		defer ml.Unlock()
	// 		if oldValue, ok := cmp.Get(key); ok {
	// 			cmp.Set(key, oldValue.(int) + 1)
	// 		} else {
	// 			cmp.Set(key, 1)
	// 		}
	// 	}()
	// }

	// wg.Wait()
	// value, _ := cmp.Get(key)
	// fmt.Println(value)
}


// 自行实现并发读写的map，借鉴ConcurrentHashMap，使用读写锁
type ConcurrentHashMap[T comparable] struct { // map支持的key类型就是comparable的类型
	mps 	[]map[T]any 	// 由多个小map构成
	seg		int 			// 小map的个数
	locks	[]sync.RWMutex 	// 每个小map配置一把读写锁。避免全局只有一把锁，影响性能
	seed 	uint32			// 每次执行farmhash统一seed
}


// cap预估map中容纳多少元素
func NewConcurrentHashMap[T comparable](seg, cap int) *ConcurrentHashMap[T] {
	mps := make([]map[T]any, seg) 
	locks := make([]sync.RWMutex, seg)

	// 每个小map开辟多大空间
	for i := 0; i < seg; i++ {
		mps[i] = make(map[T]any, cap/seg)
		locks[i] = sync.RWMutex{}
	}

	return &ConcurrentHashMap[T]{
		mps: mps,
		seg: seg,
		seed: 0,
		locks: locks,
	}
}

// getSegIndex 根据key地址获取对应的小Map
func (m *ConcurrentHashMap[T]) getSegIndex(key T) int {
	hash := int(farmhash.Hash32WithSeed(Int2Bytes(Pointer2Int(&key)), m.seed))
	return hash % m.seg
}


// 指针转int, 因为golang正常情况下不允许进行指针运算
func Pointer2Int[T comparable] (p *T) int {
	return *(*int)(unsafe.Pointer(p)) // Pointer类型可以指向任意类型, 将其转换成*int类型，再将指针进行解析，得到int
}


// 写入key, value
func (m * ConcurrentHashMap[T]) Set(key T, value any) {
	index := m.getSegIndex(key) // 根据key获取对应的小Map
	m.locks[index].Lock()  // 写锁
	defer m.locks[index].Unlock()
	m.mps[index][key] = value
}

// 根据Key读取value
func (m *ConcurrentHashMap[T]) Get(key T) (any, bool) {
	index := m.getSegIndex(key)
	m.locks[index].RLock()
	defer m.locks[index].RUnlock()
	value, ok := m.mps[index][key]
	return value, ok
}



// 自行实现并发读写的map，借鉴ConcurrentHashMap，使用自选锁
type ConcurrentHashMapSpin[T comparable] struct { // map支持的key类型就是comparable的类型
	mps 	[]map[T]any 	// 由多个小map构成
	seg		int 			// 小map的个数
	locks	[]sync.Locker 	// 每个小map配置一把读写锁。避免全局只有一把锁，影响性能
	seed 	uint32			// 每次执行farmhash统一seed
}

// cap预估map中容纳多少元素
func NewConcurrentHashMapSpin[T comparable](seg, cap int) *ConcurrentHashMapSpin[T] {
	mps := make([]map[T]any, seg) 
	locks := make([]sync.Locker, seg)

	// 每个小map开辟多大空间
	for i := 0; i < seg; i++ {
		mps[i] = make(map[T]any, cap/seg)
		locks[i] = mylock.NewSpinLock()
	}

	return &ConcurrentHashMapSpin[T]{
		mps: mps,
		seg: seg,
		seed: 0,
		locks: locks,
	}
}

// getSegIndex 根据key地址获取对应的小Map
func (m *ConcurrentHashMapSpin[T]) getSegIndexS(key T) int {
	hash := int(farmhash.Hash32WithSeed(Int2Bytes(Pointer2Int(&key)), m.seed))
	return hash % m.seg
}


// 指针转int, 因为golang正常情况下不允许进行指针运算
func Pointer2IntS[T comparable] (p *T) int {
	return *(*int)(unsafe.Pointer(p)) // Pointer类型可以指向任意类型, 将其转换成*int类型，再将指针进行解析，得到int
}


// 写入key, value
func (m * ConcurrentHashMapSpin[T]) SetS(key T, value any) {
	index := m.getSegIndexS(key) // 根据key获取对应的小Map
	m.locks[index].Lock()  
	defer m.locks[index].Unlock()
	m.mps[index][key] = value
}

// 根据Key读取value
func (m *ConcurrentHashMapSpin[T]) GetS(key T) (any, bool) {
	index := m.getSegIndexS(key)
	m.locks[index].Lock()
	defer m.locks[index].Unlock()
	value, ok := m.mps[index][key]
	return value, ok
}