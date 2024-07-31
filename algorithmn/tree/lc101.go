package tree

/*
二叉树是否对称
*/

func isSymmetric(root *BinaryTreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.LeftNode, root.RightNode)
}

func compare(left *BinaryTreeNode, right *BinaryTreeNode) bool {
	if left == nil && right == nil {
		return true 
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left.Data != right.Data {
		return false
	}

	outside := compare(left.LeftNode, right.RightNode)
	inside := compare(left.RightNode, right.LeftNode)

	return outside && inside
	
}