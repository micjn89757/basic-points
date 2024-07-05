package search

import "testing"

type lc367Example struct {
	input  int
	expect bool
}

var lc367Test = map[string]lc367Example{
	"example1": {
		input:  16,
		expect: true,
	},
	"example2": {
		input:  14,
		expect: false,
	},
}

func TestIsPerfectSquare(t *testing.T) {
	for name, te := range lc367Test {
		t.Run(name, func(t *testing.T) {
			res := isPerfectSquare(te.input)
			if res != te.expect {
				t.Errorf("expect: %t, got: %t", te.expect, res)
			}
		})
	}
}
