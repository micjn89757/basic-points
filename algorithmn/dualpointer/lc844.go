package dualpointer

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

// 注意：如果对空文本输入退格字符，文本继续为空。

// 普通解法使用栈

func backspaceCompare(s string, t string) bool {
	sIndex := len(s) - 1
	tIndex := len(t) - 1

	skipS, skipT := 0, 0 //NOTE: 表示当前待删除字符数量

	for sIndex >= 0 || tIndex >= 0 {
		// NOTE: 寻找字符串中第一个未被删除的字符
		for sIndex >= 0 {
			if s[sIndex] == '#' {
				skipS++
				sIndex--
			} else if skipS > 0 {
				skipS--
				sIndex--
			} else {
				break
			}
		}

		for tIndex >= 0 {
			if t[tIndex] == '#' {
				skipT++
				tIndex--
			} else if skipT > 0 {
				skipT--
				tIndex--
			} else {
				break
			}
		}

		// NOTE: 注意if else if 顺序
		if sIndex >= 0 && tIndex >= 0 {
			if s[sIndex] != t[tIndex] {
				return false
			}
		} else if sIndex >= 0 || tIndex >= 0 {
			return false
		}

		sIndex--
		tIndex--
	}

	return true
}
