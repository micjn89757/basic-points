package array

import (
	"testing"
)

type lc209Example struct {
	input  []int
	target int
	expect int
}

var lc209Test = map[string]lc209Example{
	"example1": {
		input:  []int{2, 3, 1, 2, 4, 3},
		target: 7,
		expect: 2,
	},
	"example2": {
		input:  []int{1, 4, 4},
		target: 4,
		expect: 1,
	},
	"example3": {
		input:  []int{1, 1, 1, 1, 1, 1, 1, 1},
		target: 11,
		expect: 0,
	},
	"example4": {
		input:  []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
		target: 80,
		expect: 6,
	},
}

func TestMinSubArrayLen(t *testing.T) {
	for name, te := range lc209Test {
		t.Run(name, func(t *testing.T) {
			res := minSubArrayLen(te.target, te.input)
			if res != te.expect {
				t.Errorf("expect: %d, got: %d", te.expect, res)
			}
		})
	}
}

