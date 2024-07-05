package array

// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其总和大于等于 target 的长度最小的
// 子数组

// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {
	res := 0

	slowIndex := 0

	sum := 0
	count := 0

	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		sum += nums[fastIndex]
		count++

		if sum >= target {
			if count < res || res == 0 {
				res = count
			}

			// NOTE: 滑动窗口起始位置回退
			for slowIndex < fastIndex {
				sum -= nums[slowIndex] //NOTE: sum有变动
				count--
				slowIndex++

				if sum >= target && count < res {
					res = count
				} else if sum < target {
					break // NOTE: 注意推出循环的条件
				}
			}

		}

	}

	return res
}
