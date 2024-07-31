package concurrency

/*
并发安全的slice
TODO
*/

type Sig uint8

const (
	add Sig = iota 
	sub 
)

type SyncSlice[T any] struct{
	ch chan Sig
	d chan T
	data []T
}


func NewSlice[T any](length int, cap int) *SyncSlice[T] {
	return &SyncSlice[T]{
		ch: make(chan Sig, 3),
		d: make(chan T, 3),
		data: make([]T, length, cap),
	}
}

func (s *SyncSlice[T]) handler() {
	select {
	case sig := <- s.ch:
		switch sig {
		case add:
			s.data = append(s.data, )
		}
	}
}

func (s *SyncSlice[T]) Add(ele T) {
	s.ch <- add
}