package array

import (
	"testing"
)

type lc59Example struct {
	input  int
	expect [][]int
}

var lc59Test = map[string]lc59Example{
	"example1": {
		input: 3,
		expect: [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
	"example2": {
		input: 1,
		expect: [][]int{
			{1},
		},
	},
}

func TestGenerateMatrix(t *testing.T) {
	for name, te := range lc59Test {
		t.Run(name, func(t *testing.T) {
			res := generateMatrix(te.input)
			for i, v1 := range res {
				for j, v2 := range v1 {
					if v2 != te.expect[i][j] {
						t.Errorf("expect: %v, got: %v", te.expect, res)
					}
				}
			}
		})
	}
}
