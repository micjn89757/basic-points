package array

import "testing"

type lc844Example struct {
	input1 string
	input2 string
	expect bool
}

var lc844Test = map[string]lc844Example{
	"example1": {
		input1: "ab#c",
		input2: "ad#c",
		expect: true,
	},
	"example2": {
		input1: "ab##",
		input2: "c#d#",
		expect: true,
	},
	"example3": {
		input1: "a#c",
		input2: "b",
		expect: false,
	},
}

func TestBackspaceCompare(t *testing.T) {
	for name, te := range lc844Test {
		t.Run(name, func(t *testing.T) {
			res := backspaceCompare(te.input1, te.input2)

			if res != te.expect {
				t.Errorf("expect: %t, got: %t", te.expect, res)
			}
		})
	}
}
