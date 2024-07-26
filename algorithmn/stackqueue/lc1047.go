package stackqueue

/*
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
*/

func removeDuplicates(s string) string {
    stack := []rune{}

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		val := stack[len(stack) - 1]
		if val != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack) - 1]
		}
	}
    
	return string(stack)
}