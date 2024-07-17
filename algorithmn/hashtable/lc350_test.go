package hashtable

import (
	"reflect"
	"testing"
)

type lc350Example struct {
	nums1 []int
	nums2 []int
	expect []int
}


var lc350Test = map[string]lc350Example{
	"example1": {
		nums1: []int{1, 2, 2, 1},
		nums2: []int{2, 2},
		expect: []int{2, 2},
	},
	"example2": {
		nums1: []int{4, 9, 5},
		nums2: []int{9, 4, 9, 8, 4},
		expect: []int{4, 9},
	},
}


// TODO
func TestIntersect(t *testing.T) {
	for name, te := range lc350Test {
		t.Run(name, func(t *testing.T) {
			res := intersect(te.nums1, te.nums2)

			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}