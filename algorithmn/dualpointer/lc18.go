package dualpointer



func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	// 排序
	sort(nums)


	for i := 0; i < len(nums); i++ {
		// 剪枝处理
		if nums[i] > target && (nums[i] > 0 || target > 0) {
			break
		}

		// 对nums[i]去重
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for k := i + 1; k < len(nums); k++ {
			// 2级剪枝
			if nums[k] + nums[i] > target && nums[k] + nums[i] >= 0 {
				break
			}

			// 对nums[k] 去重
			if k > i + 1 && nums[k] == nums[k - 1] {
				continue
			}

			left := i + 1
			right := len(nums) - 1

			for right > left {
				if nums[i] + nums[k] + nums[left] + nums[right] > target {
					right--
				} else if nums[i] + nums[k] + nums[left] + nums[right] < target {
					left ++
				} else {
					res = append(res, []int{ nums[i], nums[k], nums[left], nums[right]})

					for right > left && nums[right] == nums[right - 1] {
						right--
					}

					for right > left && nums[left] == nums[left + 1] {
						left++
					}

					right--
					left++
				}
			}
		}
	}

	return res
}
