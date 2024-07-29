package tree

import (
	"fmt"
)

type BinaryTreeNode struct {
	Data      int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

// 迭代实现（使用栈来模拟递归）
// 前序，中，左，右
func PreOrderTraversal(root *BinaryTreeNode) {
	stack := make([]*BinaryTreeNode, 0)

	if root == nil {
		return
	}

	// 加入根节点
	stack = append(stack, root)

	for len(stack) > 0 {
		// 打印并Pop
		node := stack[len(stack) - 1]
		fmt.Println(node.Data)
		stack = stack[:len(stack) - 1]

		if node.RightNode != nil {
			stack = append(stack, node.RightNode)
		}

		if node.LeftNode != nil {
			stack = append(stack, node.LeftNode)
		}
	}
}

// 二叉树中序遍历 左中右
func MiddleOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return 
	}

	stack := make([]*BinaryTreeNode, 0)

	cur := root 

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur.LeftNode)
			cur = cur.LeftNode
		} else {
			cur = stack[len(stack) - 1] // 弹出数据
			stack = stack[:len(stack) - 1]
			// 进行处理
			// ...

			cur = cur.RightNode
		}
	}
}

// 后续遍历 左右中  = 反转中右左
func PostOrderTraversal(root *BinaryTreeNode) {

}

// 递归实现
// 二叉树前序遍历 中，左，右
func RecurPreOrderTraversal(cur *BinaryTreeNode, res *[]int) {
	if cur == nil { // 遍历到叶子节点
		return
	}

	*res = append(*res, cur.Data)
	RecurPreOrderTraversal(cur.LeftNode, res)
	RecurPreOrderTraversal(cur.RightNode, res)
}

// 中序遍历 左，中，右
func RecurMiddleOrderTraversal(cur *BinaryTreeNode, res *[]int) {
	if cur == nil {
		return
	}

	RecurMiddleOrderTraversal(cur.LeftNode, res)
	*res = append(*res, cur.Data)
	RecurMiddleOrderTraversal(cur.RightNode, res)
}

// 后序遍历 左，右，中
func RecurPostOrderTraversal(cur *BinaryTreeNode, res *[]int) {
	if cur == nil {
		return
	}

	RecurMiddleOrderTraversal(cur.LeftNode, res)
	RecurMiddleOrderTraversal(cur.RightNode, res)
	*res = append(*res, cur.Data)
}


// 层序遍历，使用队列
func levelOrder(root *BinaryTreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	queue := make([]*BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmp := make([]int, 0)
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node.Data)

			if node.LeftNode != nil {
				queue = append(queue, node.LeftNode)
			}

			if node.RightNode != nil {
				queue = append(queue, node.RightNode)
			}

		}

		res = append(res, tmp)
	} 

	return res
}

// 层序遍历（递归）
func RecurLevelOrder(root *BinaryTreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *BinaryTreeNode, depth int)

	order = func(root *BinaryTreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {	// 每一层创建[]int, 这里是为了防止左右某一个节点是空的，另一个不为空，创建一次即可
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Data)

		order(root.LeftNode, depth+1)
		order(root.RightNode, depth+1)
	}

	order(root, depth)

	return arr
}

