package linkedlist

import (
	"testing"
)

type lc24Example struct {
	input []int
	expect []int
}


var lc24Test = map[string]lc24Example {
	"example1": {
		input: []int{1, 2, 3, 4},
		expect: []int{2, 1, 4, 3},
	}, 
	"example2": {
		input: []int{},
		expect: []int{},
	},
	"example3": {
		input: []int{1},
		expect: []int{1},
	},
}


func TestSwapPairs(t *testing.T) {
	for name, te := range lc24Test {
		t.Run(name, func(t *testing.T) {
			res := swapPairs(GenerateLc24LinkedList(te.input))
			expect := GenerateLc24LinkedList(te.expect)

			for res != nil && expect != nil {
				if res.Val != expect.Val {
					t.Errorf("expect: %d, got: %d", expect.Val, res.Val)
				}
				res = res.Next
				expect = expect.Next
			}

			if (res != nil && expect == nil) || (res == nil && expect != nil) {
				t.Error("error length")
			}
		})
	}
}