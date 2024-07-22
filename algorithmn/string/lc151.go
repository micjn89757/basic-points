package string

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/

func reverseWords(s string) string {
	str := []byte(s)

	// 移除前面、中间和后面存在的多余空格 (双指针法)
	slow := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if slow != 0 { // 手动控制空格，给单词之间添加空格。slow != 0说明不是第一个单词，需要在单词前添加空格。
				str[slow] = ' '
				slow++
			}

			for i < len(str) && str[i] != ' ' { // 补上该单词，遇到空格说明单词结束。
				str[slow] = str[i]
				slow++
				i++
			}
		}
	}

	str = str[0:slow] // 移除后面空格

	// 全部翻转
	reverse(str, 0, len(str)-1)

	// 反转每个单词
	last := 0
	// 注意判断条件
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			reverse(str[last:i], 0, len(str[last:i]) - 1) // 注意这里
			last = i + 1
		}
	}

	return string(str)
}
