package dualpointer

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

!有无限循环就要思考哈希表去记录
*/


// 另一种思路使用快慢指针，使用 “快慢指针” 思想，找出循环：“快指针” 每次走两步，“慢指针” 每次走一步，当二者相等时，即为一个循环周期。此时，判断是不是因为 1 引起的循环，是的话就是快乐数，否则不是快乐数。
// 注意：此题不建议用集合记录每次的计算结果来判断是否进入循环，因为这个集合可能大到无法存储；另外，也不建议使用递归，同理，如果递归层次较深，会直接导致调用栈崩溃。不要因为这个题目给出的整数是 int 型而投机取巧。

func isHappy(n int) bool {
	slow, fast := n, n 

	for {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)

		// 条件放到后面
		if slow == fast {
			break
		}
	}

	return slow == 1
}


// 计算每位之和
func bitSquareSum(n int) int {
	sum := 0 
	
	for n > 0 {
		bit := n % 10
		sum = bit * bit 
		n = n / 10
	}


	return sum
}