package designmodel

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var wg sync.WaitGroup

	for i := range 3 {
		t.Log(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstance()

		}()
	}

	wg.Wait()
	t.Log(instance)
}

func TestGetInstance2(t *testing.T) {
	var wg sync.WaitGroup

	for i := range 3 {
		t.Log(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstanceFunc()
		}()
	}

	wg.Wait()
	t.Log(instance)
}
