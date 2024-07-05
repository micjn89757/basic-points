package search


// 二分查找 easy
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// 提示：
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间
// nums 的每个元素都将在 [-9999, 9999]之间

// search 注意使用二分查找的前提必须是有序的列表且元素不重复
func search(nums []int, target int) int {
	left := 0 
	right := len(nums)

	for left < right {
		med := (left + right) / 2 // 查找中间位置

		if target == nums[med] {
			return med
		} else if target > nums[med] {
			left = med + 1
		} else {
			right = med
		}
	}

	return -1
}