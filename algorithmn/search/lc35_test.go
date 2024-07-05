package search 

import "testing"


type lc35Example struct {
	input []int 
	target int
	expect int 
}


var lc35Test = map[string]lc35Example {
	"example1": {
		input: []int{1, 3, 5,6,},
		target: 5,
		expect: 2,
	},
	"example2": {
		input: []int{1, 3, 5, 6},
		target: 2,
		expect: 1,
	},
	"example3": {
		input: []int{1, 3, 5, 6},
		target: 7,
		expect: 4,
	},
}

func TestSearchInsert(t *testing.T) {
	for name, te := range lc35Test {
		t.Run(name, func(t *testing.T) {
			res := searchInsert(te.input, te.target)
			if res != te.expect {
				t.Errorf("expect %d, got %d", te.expect, res)
			}
		})
	}
}