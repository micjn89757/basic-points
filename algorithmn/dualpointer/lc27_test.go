package array

import (
	"reflect"
	"testing"
)

type lc27Example struct {
	input  []int
	target int
	expect []int
}

var lc27Test = map[string]lc27Example{
	"example1": {
		input:  []int{3, 2, 2, 3},
		target: 3,
		expect: []int{2, 2},
	},
	"example2": {
		input:  []int{0, 1, 2, 2, 3, 0, 4, 2},
		target: 2,
		expect: []int{0, 0, 1, 3, 4},
	},
}

// sortDemo sort the top k elements of resnum
func sortDemo(nums []int, k int) {
	// NOTE: 排序次数
	for i := k - 1; i > 0; i-- {
		// NOTE: 排序流程
		for j := 0; j < k-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
			}
		}
	}
}

func TestSort(t *testing.T) {
	arr := []int{0, 2, 5, 4, 1}
	expect := []int{0, 1, 2, 4, 5}

	sortDemo(arr, len(arr))

	if !reflect.DeepEqual(arr, expect) {
		t.Errorf("expect:%v, got: %v", expect, arr)
	}
}

func TestRemoveElement(t *testing.T) {
	for name, te := range lc27Test {
		t.Run(name, func(t *testing.T) {
			resLen := removeElement(te.input, te.target)
			if resLen != len(te.expect) {
				t.Errorf("expect len: %d, got: %d", len(te.expect), resLen)
			}

			sortDemo(te.input, resLen)

			// NOTE: 比较前k个元素
			for i := 0; i < resLen; i++ {
				if te.input[i] != te.expect[i] {
					t.Errorf("expect: %v, got: %v", te.expect, te.input)
					break
				}
			}

		})
	}
}

func TestRemoveElement2(t *testing.T) {
	for name, te := range lc27Test {
		t.Run(name, func(t *testing.T) {
			resLen := removeElement2(te.input, te.target)
			if resLen != len(te.expect) {
				t.Errorf("expect len: %d, got: %d", len(te.expect), resLen)
			}

			sortDemo(te.input, resLen)

			// NOTE: 比较前k个元素
			for i := 0; i < resLen; i++ {
				if te.input[i] != te.expect[i] {
					t.Errorf("expect: %v, got: %v", te.expect, te.input)
					break
				}
			}

		})
	}
}
