package search

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

// NOTE: x平方根整数部分ans时满足k^2 <= x的最大k值, 转换成二分查找右边界
func mySqrt(x int) int {
	// 下界0， 上界x
	left := 0
	right := x

	//NOTE: 注意两种特殊情况, 因为right取不到
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	for left < right {
		mid := (left + right) >> 1

		if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}
