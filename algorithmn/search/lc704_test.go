package search

import "testing"

type lc704Example struct {
	input  []int
	target int
	expect int
}

var lc704Tests = map[string]lc704Example{
	"example1": {
		input:  []int{-1, 0, 3, 5, 9, 12},
		target: 9,
		expect: 4,
	},
	"example2": {
		input:  []int{-1, 0, 3, 5, 9, 12},
		target: 2,
		expect: -1,
	},
}

func TestSearch(t *testing.T) {
	for name, te := range lc704Tests {
		t.Run(name, func(t *testing.T) {
			res := search(te.input, te.target)
			if res != te.expect {
				t.Errorf("expect: %d, got: %d", te.expect, res)
			}
		})
	}
}
