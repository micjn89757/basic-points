package search

import (
	"reflect"
	"testing"
)

type lc34Example struct {
	input  []int
	target int
	expect []int
}

var lc34Test = map[string]lc34Example{
	"example1": {
		input:  []int{5, 7, 7, 8, 8, 10},
		target: 8,
		expect: []int{3, 4},
	},
	"example2": {
		input:  []int{5, 7, 7, 8, 8, 10},
		target: 6,
		expect: []int{-1, -1},
	},
	"example3": {
		input:  []int{},
		target: 0,
		expect: []int{-1, -1},
	},
	"example4": {
		input:  []int{1},
		target: 0,
		expect: []int{-1, -1},
	},
}

func TestSearchRange(t *testing.T) {
	for name, te := range lc34Test {
		t.Run(name, func(t *testing.T) {
			res := searchRange(te.input, te.target)
			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}
