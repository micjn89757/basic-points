package hashtable

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

// 根据题目要求，我们需要在字符串 s 寻找字符串 p 的异位词。因为字符串 p 的异位词的长度一定与字符串 p 的长度相同，所以我们可以在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；当窗口中每种字母的数量与字符串 p 中每种字母的数量相同时，则说明当前窗口为字符串 p 的异位词。


func findAnagrams(s, p string) []int {
	sLen, pLen := len(s), len(p)
    ans := make([]int, 0)
    if sLen < pLen {
        return nil
    }

    scnt, pcnt := [26]uint16{}, [26]uint16{} // 记录各个字母的数量(scnt表示子串中字符数量，pcnt表示target中的数量)

    // 初始化窗口（第一个窗口）
    for i := 0; i < pLen; i++ {
        scnt[s[i]-'a']++
        pcnt[p[i]-'a']++
    }

    if scnt == pcnt { // 数组可比较,是否有同等数量的字母
        ans = append(ans, 0)
    }

    // 滑动窗口
    for i := 1; i <= sLen-pLen; i++ {
        // 每移动一次将滑动窗口左边界元素减少，右边增加
        scnt[s[i-1] - 'a']--
        scnt[s[i+pLen-1] - 'a']++
        if scnt == pcnt {
            ans = append(ans, i)
        }
    }

    return ans
}