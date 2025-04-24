package cache

import (
	// "fmt"
	"slices"
	"testing"
)

func TestXxx(t *testing.T) {
	slice := make([]int, 2)

	slice[0] = 1
	slice[1] = 2

	t.Log(slice[:0])
	t.Log(slice[2:])
}

func TestCombine(t *testing.T) {
	t.Log(combine(4, 2))
}

func combine(n int, k int) [][]int {
    ans := [][]int{}

    path := []int{}

    var dfs func(int)
    dfs = func(i int) {  
        if i == k {
            ans = append(ans, slices.Clone(path))
            return 
        }

        for j := i; j < n; j++ {
            path = append(path, j + 1)
            dfs(j + 1)
            path = path[:len(path) - 1]
        }
    }

    dfs(0)

    return ans
}