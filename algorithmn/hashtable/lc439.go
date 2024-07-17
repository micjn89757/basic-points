package hashtable

func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    tmp := make(map[int]int, 0)

    for _, v := range nums1 {
        tmp[v] = 1
    }


    for _, v := range nums2 {
        if _, ok := tmp[v]; ok {
            tmp[v]++
        }
    }

    for k, v := range tmp {
        if v > 1 {
            res = append(res, k)
        }
    }

    return res
}