package tree


func findBottomLeftValue(root *BinaryTreeNode) int {
    if root == nil {
        return 0
    }

    queue := make([]*BinaryTreeNode, 0)

    res := 0

    queue = append(queue, root)

    for len(queue) > 0 {
        length := len(queue)
        
        for i := 0; i < length; i++ {
            node := queue[0]
            if i == 0 {
                res = node.Data
            }        
            queue = queue[1:]

            

            // 加入节点
            if node.LeftNode != nil {
                queue = append(queue, node.LeftNode)
            }

            if node.RightNode != nil {
                queue = append(queue, node.RightNode)
            }
        }

        
    }

    return res
}