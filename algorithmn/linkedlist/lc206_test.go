package linkedlist

import (
	"testing"
)


type lc206Example struct {
	input []int 
	expect []int
}


var lc206Test = map[string]lc206Example {
	"example1": {
		input: []int{1, 2, 3, 4, 5},
		expect: []int{5, 4, 3, 2, 1},
	},
	"example2": {
		input: []int{1, 2},
		expect: []int{2, 1},
	},
	"example3": {
		input: []int{},
		expect: []int{},
	},
}


func TestReverseList(t *testing.T) {
	for name, te := range lc206Test {
		t.Run(name, func(t *testing.T) {
			res := GenerateLc206LinkedList(te.input)
			res = reverseList(res)
			exp := GenerateLc206LinkedList(te.expect)

			for res != nil && exp != nil {
				if res.Val != exp.Val {
					t.Errorf("expect: %d, got: %d", exp.Val, res.Val)
				}
				res = res.Next
				exp = exp.Next
			}

			if ( res == nil && exp != nil ) || ( res !=nil && exp == nil) {
				t.Error("res len error")
			}
		})
	}
}