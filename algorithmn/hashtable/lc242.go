package hashtable


// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

func isAnagram(s string, t string) bool {
	record := make([]int, 26)

    for _, ch := range s {
        record[ch-rune('a')]++
    }

    for _, ch := range t {
        if record[ch-rune('a')] > 0 {
            record[ch-rune('a')]--
        } else {
            record[ch-rune('a')]++
        }
    }

    for _, v := range record {
        if v > 0{
            return false
        }
    }

    return true
}