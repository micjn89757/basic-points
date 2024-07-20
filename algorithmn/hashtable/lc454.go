package hashtable

/*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/


func fourSumCount(nums1 []int, nums2 []int, nums3 []int ,nums4 []int) int {
	resMap := make(map[int]int)	// key放两数之和, value放两数之和出现的次数
	count := 0 // 结果

	for _, vi := range nums1 {
		for _, vj := range nums2 {
			sum := vi + vj
			// map中key不存在则value返回对应的零值
			if resMap[sum] != 0  {
				resMap[sum]++
			} else {
				resMap[sum] = 1
			}
		}
	}


	for _, vi := range nums3 {
		for _, vj := range nums4 {
			sum := vi + vj
			if resMap[0 - sum] != 0 {
				count += resMap[0 - sum]
			}
		}
	}


	return count
}