package dualpointer

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	left := 0
	right := len(nums) - 1

	resRight := len(res) - 1

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[resRight] = nums[left] * nums[left]
			left++
		} else {
			res[resRight] = nums[right] * nums[right]
			right--
		}
		resRight--
	}

	return res
}
