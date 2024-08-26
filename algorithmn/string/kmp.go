package string


/*
KMP解决的问题：字符串匹配
*/

func kmp(a, b string) []int {
	if len(b) == 0 {
		return nil
	}


	next := getNext(b)
	res := make([]int, 0)

	j := 0
	for i := 0; i < len(a); i++ {
		// 匹配失败
		for a[i] != b[j] && j > 0 {
			j = next[j - 1]
		}

		// 匹配成功
		if a[i] == b[j] {
			j++
		}

		// 完全匹配
		if j == len(b) {
			// 记录位置
			res = append(res, i - len(b) + 1)
			j = 0
		}
	}

	return res
}



// 构建next
func getNext(s string) []int {
	next := make([]int, len(s))
	j := 0 // j记录了当前最长公共前后缀长度，以及当前前缀的尾部
	next[0] = j


	for i := 1; i < len(s); i++ {
		// 匹配失败，回溯到相同位置
		for j > 0 && s[i] != s[j] {
			j = next[j - 1]
		}

		// 匹配成功
		if s[i] == s[j] {
			j++
		}
		next[i] = j 
	}

	return next 
}