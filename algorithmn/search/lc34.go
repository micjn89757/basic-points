package search

// 在排序数组中查找元素的第一个和最后一个位置
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

func searchRange(nums []int, target int) []int {
	leftRange := 0
	rightRange := 0

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftRange = searchLeftRange(nums, target)
	rightRange = searchRightRange(nums, target)

	return []int{leftRange, rightRange}
}

func searchLeftRange(nums []int, target int) int {
	leftRange := 0
	left := 0
	right := len(nums)

	// seek left range
	for left < right {
		mid := (left + right) >> 1
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// NOTE: the range of right
	if right >= 0 && right < len(nums) && nums[right] == target {
		leftRange = right
	} else {
		leftRange = -1
	}

	return leftRange
}

func searchRightRange(nums []int, target int) int {
	rightRange := 0
	left := 0
	right := len(nums)

	// seek right range
	for left < right {
		mid := (left + right) >> 1
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	//NOTE: the range of left - 1
	if (left-1) >= 0 && (left-1) < len(nums) && nums[left-1] == target {
		rightRange = left - 1
	} else {
		rightRange = -1
	}

	return rightRange
}
