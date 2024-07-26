package stackqueue

import (
	"strconv"
)

/*
给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

注意：

有效的算符为 '+'、'-'、'*' 和 '/' 。
每个操作数（运算对象）都可以是一个整数或者另一个表达式。
两个整数之间的除法总是 向零截断 。
表达式中不含除零运算。
输入是一个根据逆波兰表示法表示的算术表达式。
答案及所有中间计算结果可以用 32 位 整数表示。
*/

func evalRPN(tokens []string) int {
	stack := []string{}


	for _, c := range tokens {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		switch c {
		case "+":
			// 出栈两数字相加
			if len(stack) < 2 {
				return -1
			}
			n1, _:= strconv.Atoi(stack[len(stack) - 1])
			n2, _:= strconv.Atoi(stack[len(stack) - 2])
			stack = stack[:len(stack) - 2]

			// 操作结果推入栈中
			stack = append(stack, strconv.Itoa(n1 + n2))
		case "-":
			if len(stack) < 2 {
				return -1
			}
			n1, _:= strconv.Atoi(stack[len(stack) - 1])
			n2, _:= strconv.Atoi(stack[len(stack) - 2])
			stack = stack[:len(stack) - 2]

			// 操作结果推入栈中
			stack = append(stack, strconv.Itoa(n2 - n1))
		case "*":
			if len(stack) < 2 {
				return -1
			}
			n1, _:= strconv.Atoi(stack[len(stack) - 1])
			n2, _:= strconv.Atoi(stack[len(stack) - 2])
			stack = stack[:len(stack) - 2]

			// 操作结果推入栈中
			stack = append(stack, strconv.Itoa(n2 * n1))
		case "/":
			if len(stack) < 2 {
				return -1
			}
			n1, _:= strconv.Atoi(stack[len(stack) - 1])
			n2, _:= strconv.Atoi(stack[len(stack) - 2])
			stack = stack[:len(stack) - 2]

			// 操作结果推入栈中
			stack = append(stack, strconv.Itoa(n2 / n1))
		default: 	// 数字直接入栈
			stack = append(stack, c)
		}

	}


	ret, _ := strconv.Atoi(stack[0])

	return ret
}

