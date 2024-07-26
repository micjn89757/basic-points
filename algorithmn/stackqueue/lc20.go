package stackqueue

func isValid(s string) bool {
	stack := []rune{}

	// 将括号压入栈中, 如果栈顶元素和当前匹配的字符是一对，则出栈，继续匹配
	for _, c := range s {
		// 压栈
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		// 从栈顶取出元素, 与字符进行对比
		val := stack[len(stack)-1]
		var need rune
		switch val {
		case '(':
			need = ')'
		case '{':
			need = '}'
		case '[':
			need = ']'
		default: // 匹配到了右括号
			return false
		}

		if need != c {
			stack = append(stack, c)
		} else {
			// 匹配成功就出栈
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0 
}