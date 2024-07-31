package tree

/*
翻转二叉树
*/

// 递归(深度优先)
func invertTree(root *BinaryTreeNode) *BinaryTreeNode {
    if root == nil {
        return nil
    }

    reversei(root)

    return root
}

func reversei(root *BinaryTreeNode) {
    if root == nil {
        return  
    }

    root.LeftNode, root.RightNode = root.RightNode, root.LeftNode 

    reversei(root.LeftNode)
    reversei(root.RightNode)
}

