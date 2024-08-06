package string

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长 
子串的长度。
*/

func lengthOfLongestSubstring(s string) int {
    
    // val统计字符位置
    pattern := make(map[rune]int)
    res := 0
    left := 0
    str := []rune(s)

    for right := 0; right < len(str); right++ {
        if index, ok := pattern[str[right]]; ok && index >= left {
            left = index + 1
        } 

        pattern[str[right]] = right

        tmp := right - left + 1
        if tmp > res {
            res = tmp
        }
    }

    return res

}



