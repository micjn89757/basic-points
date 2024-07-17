package hashtable

import "slices"
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。


// 注意到，如果把 aab,aba,baa 按照字母从小到大排序，我们可以得到同一个字符串 aab。

// 而对于每种字母出现次数不同于 aab 的字符串，例如 abb 和 bab，排序后为 abb，不等于 aab。

// 所以当且仅当两个字符串排序后一样，这两个字符串才能分到同一组。

// 根据这一点，我们可以用哈希表来分组，把排序后的字符串当作 key，原字符串组成的列表（即答案）当作 value。

// 最后把所有 value 加到一个列表中返回。


func groupAnagrams(strs []string) [][]string {
	    m := map[string][]string{}
    for _, s := range strs {
        t := []byte(s)
        slices.Sort(t)
        sortedS := string(t)  // 排好序的字符串当作key
        m[sortedS] = append(m[sortedS], s)
    }

     ans := make([][]string, 0, len(m))
     for _, v := range m {  // 将分好组的加入到结果中
        ans = append(ans, v)
     }

     return ans
}