package tree

/*
	是否是平衡二叉树
*/

func isBalanced(root *BinaryTreeNode) bool {
    h := getHeight(root)
    if h == -1 {
        return false
    }
    return true
}
// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
func getHeight(root *BinaryTreeNode) int {
    if root == nil {
        return 0
    }
    l, r := getHeight(root.LeftNode), getHeight(root.RightNode)
    if l == -1 || r == -1 {
        return -1
    }
    if l - r > 1 || r - l > 1 {
        return -1
    }
    return max(l, r) + 1
}

