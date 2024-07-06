package array

func generateMatrix(n int) [][]int {
	// 定义一个二维矩阵
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 定义每循环一个圈的起始位置
	startx, starty := 0, 0

	// 矩阵中间的位置，横纵坐标相同
	mid := n / 2

	// 每一个空格赋值
	count := 1

	// 控制每一圈里每一条边遍历的长度, 每次循环右边界收缩一位
	offset := 1

	// 当前坐标
	i, j := 0, 0

	// 循环几圈，如果是奇数会多出一个矩阵中间值需要单独处理
	for loop := n / 2; loop > 0; loop-- {
		i = startx
		j = starty

		// 以下四个循环遵循左闭右开
		// 左上到右上
		for ; j < n-offset; j++ {
			res[i][j] = count
			count++
		}

		// 右上到右下
		for ; i < n-offset; i++ {
			res[i][j] = count
			count++
		}

		// 注意这里每一圈的起始位置不同
		// 右下到左下
		for ; j > startx; j-- {
			res[i][j] = count
			count++
		}

		// 左下到左上
		for ; i > starty; i-- {
			res[i][j] = count 
			count++
		}

		// 第二圈开始的时候起始位置都要+1
		startx++
		starty++

		offset++
	}

	// n为奇数，要给中间位置单独赋值
	if n % 2 != 0 {
		res[mid][mid] = count
	}

	return res
}