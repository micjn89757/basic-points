package hashtable

// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

func intersect(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    tmp := map[int]int{}


    for _, v := range nums1 {
        if tmp[v] != 0 {
            tmp[v]++
        } else {
            tmp[v] = 1
        }
    }

    for _, v := range nums2 {
        if va, ok := tmp[v]; ok && va > 0 {
            res = append(res, v)
            tmp[v]--
        }
    }

    return res 
}