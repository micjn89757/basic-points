package array

import (
	"reflect"
	"testing"
)

type lc283Example struct {
	input  []int
	expect []int
}

var lc283Test = map[string]lc283Example{
	"example1": {
		input:  []int{0, 1, 0, 3, 12},
		expect: []int{1, 3, 12, 0, 0},
	},
	"example2": {
		input:  []int{0},
		expect: []int{0},
	},
}

func TestMoveZeros(t *testing.T) {
	for name, te := range lc283Test {
		t.Run(name, func(t *testing.T) {
			moveZeros(te.input)
			if !reflect.DeepEqual(te.input, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, te.input)
			}
		})
	}
}
