package hashtable

import (
	"testing"
)


type lc242Example struct {
	s string 
	t string 
	expect bool
}

var lc242Test = map[string]lc242Example {
	"example1": {
		s: "anagram",
		t: "nagaram",
		expect: true,
	},
	"example2": {
		s: "rat",
		t: "car",
		expect: false,
	},
	"example3": {
		s: "",
		t: "b",
		expect: false,
	},
}

func TestIsAnagram(t *testing.T) {
	for name, te := range lc242Test {
		t.Run(name, func(t *testing.T) {
			res := isAnagram(te.s, te.t)

			if res != te.expect {
				t.Errorf("expect: %t, got: %t", te.expect, res)
			}
		})
	}
}


