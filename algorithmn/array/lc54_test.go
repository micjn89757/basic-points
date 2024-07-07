package array

import (
	"reflect"
	"testing"
)

type lc54Example struct {
	input  [][]int
	expect []int
}

var lc54Test = map[string]lc54Example{
	"example1": {
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	"example2": {
		input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	"example3": {
		input: [][]int {
			{1},
			{2},
			{3},
			{4},
			{5},
		},
		expect: []int{1, 2, 3, 4, 5},
	},
}

func TestSpiralOrder(t *testing.T) {
	for name, te := range lc54Test {
		t.Run(name, func(t *testing.T) {
			res := spiralOrder(te.input)

			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}