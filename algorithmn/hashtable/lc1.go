package hashtable

// 两数之和 easy
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9

// twoSum1 暴力枚举, 时间复杂度O(n^2), 空间复杂度O(1)
func twoSum1(nums []int, target int) []int {
	var res []int 

	for i, v := range nums{
		for j := i+1; j < len(nums); j++ {
			if v + nums[j] == target {
				res = []int{i, j}
			}
		}
	}

	return res 
}


// TwoSum2 使用Map， 时间复杂度O(n), 空间复杂度O(n)
func twoSum2(nums []int, target int) []int {
	var res []int
	tempMap := make(map[int]int, len(nums)) // key: 元素值， value: 元素下标
	
	for i, v := range nums {
		subNum := target - v  // target - v是否存在
		subNumIndex, ok := tempMap[subNum]

		if ok {
			res = []int{subNumIndex, i}
		}else {
			tempMap[v] = i
		}
	}
	
	return res
}