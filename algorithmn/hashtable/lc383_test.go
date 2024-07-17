package hashtable

import "testing"

type lc383Example struct {
	ransomNote	string 
	magazine 	string 
	expect 		bool
}

var lc383Test = map[string]lc383Example {
	"example1": {
		ransomNote: "a",
		magazine: "b",
		expect: false,
	},
	"example2": {
		ransomNote: "aa",
		magazine: "ab",
		expect: false,
	},
	"example3": {
		ransomNote: "aa",
		magazine: "aab",
		expect: true,
	},
}


func TestCanConstruct(t *testing.T) {
	for name, te := range lc383Test {
		t.Run(name, func(t *testing.T) {
			res := canConstruct(te.ransomNote, te.magazine)
			if res != te.expect {
				t.Errorf("expect: %t, got: %t", te.expect, res)
			}
		})
	}
}