package stackqueue

import "testing"

func TestTopK(t *testing.T) {
	nums := []int{5,2,5,3,5,3,1,1,3}
	k := 2
	t.Logf("%#v",topKFrequent(nums, k))
	// t.Logf("%v", -3 / 2)
}