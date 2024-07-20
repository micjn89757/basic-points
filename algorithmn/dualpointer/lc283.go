package dualpointer

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func moveZeros(nums []int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	// NOTE: 移动后剩余元素都是0
	for ; slowIndex < len(nums); slowIndex++ {
		nums[slowIndex] = 0
	}
}
