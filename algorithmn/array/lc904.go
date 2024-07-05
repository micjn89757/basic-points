package array

func totalFruit(fruits []int) int {
	// 每颗树上恰好摘一个水果，不同编号表示不同类别，篮子有两个，一个篮子一个类型水果不限量，从左向右移动

	res := 0                       // 最大树数量
	count := 0                     // 当前收集策略，采摘的水果树数量
	free := 2                      // 空闲框子
	hamper := make(map[int]int, 2) // key表示种类， value表示水果数量

	slowIndex := 0

	for fastIndex := 0; fastIndex < len(fruits); fastIndex++ {
		_, ok := hamper[fruits[fastIndex]]
		count++
		// 没有空闲框子，并且此种类不存在，停止收集
		if free == 0 && !ok {
			for free < 1 {
				// 开始移动采摘起点, 这里需要对框子进行处理
				hamper[fruits[slowIndex]] -= 1
				// 如果为0，释放框子
				if hamper[fruits[slowIndex]] == 0 {
					delete(hamper, fruits[slowIndex])
					free++
				}
				slowIndex++
				count--
			}

		}

		// 其他情况都可以收集
		_, ok = hamper[fruits[fastIndex]]

		if ok {
			hamper[fruits[fastIndex]] += 1
		} else {
			hamper[fruits[fastIndex]] = 1
			free--
		}

		if count > res {
			res = count
		}

	}

	return res
}
