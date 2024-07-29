package tree


func rightSideView(root *BinaryTreeNode) []int {
	// 层序遍历
	if root == nil {
		return nil
	}

	queue := make([]*BinaryTreeNode, 0)
	res := make([]int, 0)

	queue = append(queue, root)
	
	for len(queue) > 0 {
		// 保存快照
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == length - 1 {
				res = append(res, node.Data)
			}

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