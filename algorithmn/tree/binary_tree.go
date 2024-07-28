package tree


type BinaryTreeNode struct {
	Data 		int 
	LeftNode 	*BinaryTreeNode 
	RightNode	*BinaryTreeNode
}

// 迭代实现（使用栈来模拟递归）

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