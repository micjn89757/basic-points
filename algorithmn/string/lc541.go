package string

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
*/

func reverseStr(s string, k int) string {
	str := []byte(s)

	for i := 0; i < len(str); i += 2 * k {
        if (i + k) <= len(str) {
            reverse(str, i, i + k - 1)
        } else {
            // 剩余字符小于k
            reverse(str, i, len(str) - 1)

        }

	}

	return string(str)
}


func reverse(str []byte, left int, right int) {
	for left < right {
		tmp := str[left]
		str[left] = str[right]
		str[right] = tmp
        left++
        right--
	}
}