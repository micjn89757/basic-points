package array

import (
	"reflect"
	"testing"
)

type lc977Example struct {
	input  []int
	expect []int
}

var lc977Test = map[string]lc977Example{
	"example1": {
		input:  []int{-4, -1, 0, 3, 10},
		expect: []int{0, 1, 9, 16, 100},
	},
	"example2": {
		input:  []int{-7, -3, 2, 3, 11},
		expect: []int{4, 9, 9, 49, 121},
	},
}

func TestSortedSquares(t *testing.T) {
	for name, te := range lc977Test {
		t.Run(name, func(t *testing.T) {
			res := sortedSquares(te.input)

			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}
