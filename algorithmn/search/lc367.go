package search

// 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

// 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

// 不能使用任何内置的库函数，如  sqrt

func isPerfectSquare(num int) bool {
	left := 0
	right := num

	//NOTE: 注意right边界值的两种特殊情况, 因为这种方式的二分查找取不到right
	if num == 1 {
		return true
	}

	if num == 0 {
		return true
	}

	for left < right {
		mid := (left + right) >> 1

		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return false
}
