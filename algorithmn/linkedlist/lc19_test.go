package linkedlist

import "testing"

type lc19Example struct {
	input []int
	target int
	expect []int
}


var lc19Test = map[string]lc19Example {
	"example1": {
		input: []int{1, 2, 3, 4, 5},
		target: 2,
		expect: []int{1, 2, 3, 5},
	}, 
	"example2": {
		input: []int{1},
		target: 1,
		expect: []int{},
	},
	"example3": {
		input: []int{1, 2},
		target: 1,
		expect: []int{1},
	},
}


func TestRemoveNthFromEnd(t *testing.T) {
	for name, te := range lc19Test {
		t.Run(name, func(t *testing.T) {
			res := removeNthFromEnd(GenerateLinkedList(te.input), te.target)
			expect := GenerateLinkedList(te.expect)

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