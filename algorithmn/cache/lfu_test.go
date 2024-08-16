package cache

import "testing"

func TestXxx(t *testing.T) {
	slice := make([]int, 2)

	slice[0] = 1
	slice[1] = 2

	t.Log(slice[:0])
	t.Log(slice[2:])
	t.Log(type(struct{}))
}