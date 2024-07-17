package hashtable

import (
	"reflect"
	"testing"
)

type lc49Example struct {
	input  []string
	expect [][]string
}

var lc49Test = map[string]lc49Example{
	"example1": {
		input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expect: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
	"example2": {
		input:  []string{""},
		expect: [][]string{{""}},
	},
	"example3": {
		input:  []string{"a"},
		expect: [][]string{{"a"}},
	},
}

// TODO
func TestGroupAnagrams(t *testing.T) {
	for name, te := range lc49Test{
		t.Run(name, func (t *testing.T) {
			res := groupAnagrams(te.input)

			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expect: %v, got: %v", te.expect, res)
			}
		})
	}
}