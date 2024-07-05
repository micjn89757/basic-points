package array

import "testing"

type lc76Example struct {
	s      string
	t      string
	expect string
}

var lc76Test = map[string]lc76Example{
	"example1": {
		s:      "ADOBECODEBANC",
		t:      "ABC",
		expect: "BANC",
	},
	"example2": {
		s:      "a",
		t:      "a",
		expect: "a",
	},
	"example3": {
		s:      "a",
		t:      "aa",
		expect: "",
	},
}

func TestMinWindow(t *testing.T) {
	for name, te := range lc76Test {
		t.Run(name, func(t *testing.T) {
			res := minWindow(te.s, te.t)

			if res != te.expect {
				t.Errorf("expect: %s, got: %s", te.expect, res)
			}
		})
	}
}
