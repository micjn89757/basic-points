package tree

/*
求树的深度
*/


// 递归
func RecurMaxDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	
	return max(RecurMaxDepth(root.LeftNode), RecurMaxDepth(root.RightNode)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}