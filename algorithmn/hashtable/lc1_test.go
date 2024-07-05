package hashtable

import (
	"reflect"
	"testing"
)

type testExample struct {
	input  []int
	target int
	expect []int
}

var tests = map[string]testExample{
	"example1": {input: []int{2, 7, 11, 15}, target: 9, expect: []int{0, 1}},
	"example2": {input: []int{3, 2, 4}, target: 6, expect: []int{1, 2}},
	"example3": {input: []int{3, 3}, target: 6, expect: []int{0, 1}},
}

func TestTwoSum1(t *testing.T) {
	for name, te := range tests {
		t.Run(name, func(t *testing.T) { // 子测试
			res := twoSum1(te.input, te.target)
			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expected:%#v, got:%#v\n", te.expect, res)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	for name, te := range tests {
		t.Run(name, func(t *testing.T) { // 子测试
			res := twoSum2(te.input, te.target)
			if !reflect.DeepEqual(res, te.expect) {
				t.Errorf("expected:%#v, got:%#v\n", te.expect, res)
			}
		})
	}
}
