package search

import "testing"

type lc69Example struct {
	input  int
	expect int
}

var lc69Test = map[string]lc69Example{
	"example1": {
		input:  4,
		expect: 2,
	},
	"example2": {
		input:  8,
		expect: 2,
	},
}

func TestMySqrt(t *testing.T) {

	for name, te := range lc69Test {
		t.Run(name, func(t *testing.T) {
			res := mySqrt(te.input)
			if res != te.expect {
				t.Errorf("expect: %d, get: %d", te.expect, res)
			}
		})
	}

}
