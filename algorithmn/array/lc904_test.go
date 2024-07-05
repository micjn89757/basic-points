package array

import (
	"testing"
)

type lc904Example struct {
	input  []int
	expect int
}

var lc907Test = map[string]lc904Example{
	"example1": {
		input:  []int{1, 2, 1},
		expect: 3,
	},
	"example2": {
		input:  []int{0, 1, 2, 2},
		expect: 3,
	},
	"example3": {
		input:  []int{1, 2, 3, 2, 2},
		expect: 4,
	},
	"example4": {
		input:  []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
		expect: 5,
	},
}

func TestTotalFruit(t *testing.T) {
	for name, te := range lc907Test {
		t.Run(name, func(t *testing.T) {
			res := totalFruit(te.input)

			if res != te.expect {
				t.Errorf("expect: %d, got: %d", te.expect, res)
			}
		})
	}
}
