package array

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

// 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k。
func removeElement(nums []int, val int) int {
	size := len(nums)
	// NOTE: 寻找val, 注意这个size是要变化的
	for i := 0; i < size; i++ {
		// NOTE: 调整
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i--    //NOTE: 下标也要向前移动，否则下标后刚移动过来的元素无法检测到
			size-- // size一定要在整理完数组后移动
		}
	}

	return size
}

// NOTE: 快慢指针
// 快指针：寻找新数组元素
// 慢指针：指向更新新数组下标的位置, 新数组长度为慢指针
func removeElement2(nums []int, val int) int {
	size := len(nums)
	slowIndex := 0

	for fastIndex := 0; fastIndex < size; fastIndex++ {
		if val != nums[fastIndex] {
			// NOTE: fastIndex找到新数组元素，使用slowIndex填充
			nums[slowIndex] = nums[fastIndex]
			slowIndex++ // NOTE: 填充完后移动到下一个位置
		}
	}

	return slowIndex
}
