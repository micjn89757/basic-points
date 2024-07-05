package search 

// 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 提示
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为 无重复元素 的 升序 排列数组
// -104 <= target <= 104

func searchInsert(nums []int, target int) int {
	left := 0 
	right := len(nums)
	
	for left < right {
		med := (left + right) >> 1

		if nums[med] == target {
			return med
		} else if nums[med] > target {
			right = med
		} else {
			left = med + 1
		}
	}

	return right
}