package hashtable

import (
	"reflect"
	"testing"
)

type lc438Example struct {
	s string 
	p string 
	expect []int
}


var lc438Test = map[string]lc438Example{
	"example1": {
		s: "cbaebabacd",
		p: "abc",
		expect: []int{0, 6},
	},
	"example2": {
		s: "abab",
		p: "ab",
		expect: []int{0, 1, 2},
	},
}


func TestFindAnagrams(t *testing.T) {
	for name, te := range lc438Test {
		t.Run(name, func(t *testing.T) {
			res := findAnagrams(te.s, te.p)
			if !reflect.DeepEqual(te.expect, res) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}