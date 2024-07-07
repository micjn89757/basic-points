package array

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素

func spiralOrder(matrix [][]int) []int {
	// 数组为空，直接返回答案
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }

    // 记录行数列数
    res := make([]int, 0)

    // 上下左右边界
    u := 0
    d := len(matrix) - 1
    l := 0
    r := len(matrix[0]) - 1

    // 向右移动到最右，此时第一行因为已经使用过了，可以将其从图中删去，体现在代码中就是重新定义上边界
    // 判断若重新定义后，上下边界交错，表明螺旋矩阵遍历结束，跳出循环，返回答案
    // 若上下边界不交错，则遍历还未结束，接着向下向左向上移动，操作过程与第一，二步同理
    // 不断循环以上步骤，直到某两条边界交错，跳出循环，返回答案
    for {
        // 左-右
        for i := l; i <= r; i++ {
            res = append(res, matrix[u][i])
        }
        // 调整上边界
        u++
        if (u > d) {
            break;
        }

        // 右-下
        for i := u; i <= d; i++ {
            res = append(res, matrix[i][r])
        }

        // 调整右边界
        r--
        if (r < l) {
            break
        }

        // 右下-左下
        for i := r; i >= l; i-- {
            res = append(res, matrix[d][i])
        }

        // 调整下边界
        d--
        if (d < u) {
            break 
        }

        // 左下-左上
        for i := d; i >= u; i-- {
            res = append(res, matrix[i][l])
        }

        // 调整左边界
        l++
        if (l > r) {
            break
        }
    }

    return res
}
