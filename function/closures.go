package function

// 闭包

func intSeq() func() int {
	i := 0 
	return func() int {
		i++
		return i
	}
}


