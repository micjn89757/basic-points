package array

import "math"

func minWindow(s string, t string) string {
	// 记录字串截取位置，左闭右开
    resLeft := -1
    resRight := -1

    // 当前最小子串长度
    resLen := math.MaxInt32
    slen := len(s)

    // 注意这里如果使用make分配会出现超时
    // 存储t中的字符
    target := map[byte]int{}
    // 遍历过程中统计字符
    temp := map[byte]int{}
    
    // 记录t
    for i := 0; i < len(t); i++ {
        target[t[i]]++
    }

    // 检查子串是否已经满足t
    check := func() bool {
        for k, v := range target {
            if temp[k] < v {
                return false
            }
        }
        return true
    }

    for left, right := 0, 0; right < slen; right++ {
        // t中存在这个字符
        if target[s[right]] > 0 {
            temp[s[right]]++
        }

        // 调动左侧位置
        for check() && left <= right {
            // 记录子串长度, 并调整起始点位置
            if (right - left + 1 < resLen) {
                resLen = right - left + 1
                resLeft, resRight = left, left + resLen
            }

            if _, ok := target[s[left]]; ok {
                temp[s[left]]--
            }

            left++
        }
    }

    if resLeft == -1 {
        return ""
    }

    return s[resLeft:resRight]
}
