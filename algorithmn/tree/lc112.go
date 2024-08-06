package tree 


/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。
*/


func hasPathSum(root *BinaryTreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    var traversal func(cur *BinaryTreeNode, count int) bool
    traversal = func(cur *BinaryTreeNode, count int) bool {
        if cur.LeftNode == nil && cur.RightNode == nil && count == 0 {  // 叶子节点并且计数器为0，找到路径
            return true 
        }

        if cur.LeftNode == nil && cur.RightNode == nil {    // 叶子节点，但是不是这个路径
            return false
        }

        if cur.LeftNode != nil {    // 左子树回溯
            count -= cur.LeftNode.Data
            lRes := traversal(cur.LeftNode, count)
            if lRes == true {
                return true
            }
            count += cur.LeftNode.Data
        }

        if cur.RightNode != nil {   // 右子树回溯
            count -= cur.RightNode.Data 
            rRes := traversal(cur.RightNode, count)
            if rRes == true {
                return true
            }
            count += cur.RightNode.Data
        }

        return false    // 左右子树都没有满足条件的
    }

    res := traversal(root, targetSum - root.Data)

    return res
}
