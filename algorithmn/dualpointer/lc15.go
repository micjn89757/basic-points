package dualpointer

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort(nums)

	// i 表示结果中最小的数
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}

		// 找的是在[i, right] 之间的三个数

		// 去重
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for right > left {
			if (nums[i] + nums[left] + nums[right]) > 0 {
				right--
			} else if (nums[i] + nums[left] + nums[right]) < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 找到一个三元组后对left和right去重, 防止循环结束之前找到相同的b, c
				for right > left && nums[right] == nums[right - 1] {
					right--
				}

				for right > left && nums[left] == nums[left + 1] {
					left++
				}


				// 找到答案双指针同时收缩
				right--
				left++
			}
		}
	}

	return res
}


// 排序
func sort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j + 1] {
				tmp := nums[j + 1]
				nums[j + 1] = nums[j]
				nums[j] = tmp
			}
		}
	}
}