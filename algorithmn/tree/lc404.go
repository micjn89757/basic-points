package tree


/*
给定二叉树的根节点 root ，返回所有左叶子之和
*/


func sumOfLeftLeaves(root *BinaryTreeNode) int {
    if root == nil {
        return 0
    }

    // 左子树的左节点和右子树的左叶子节点
    left := sumOfLeftLeaves(root.LeftNode)
    if root.LeftNode != nil && root.LeftNode.LeftNode == nil && root.LeftNode.RightNode == nil { // 只需要左叶子节点
        left = root.LeftNode.Data
    }

    right := sumOfLeftLeaves(root.RightNode)
    
    return left + right
}