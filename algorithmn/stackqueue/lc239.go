package stackqueue


/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
*/

func maxSlidingWindow(nums []int, k int) []int {
	// 单调队列， 保证最左边是当前最大的，依次递减
    var stack []int // 存放 <= i 中的
    var res []int
    for i, v := range nums {
        
        // 如果新入队元素大于队内所有元素，则全部剔除
        for len(stack) > 0 && v >= nums[stack[len(stack)-1]] { 
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, i)

        if i-k+1 > stack[0] {   // 判断滑动窗口队头元素的下标如果大于stack[0]，说明它过期了
            stack = stack[1:]
        }

        if i+1 >= k {       // 从第一个窗口开始每次添加最大值坐标
            res = append(res, nums[stack[0]])   // stack[0] 放的是最大值坐标
        }   
    }
    return res
}