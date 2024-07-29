package tree



func averageOfLevels(root *BinaryTreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)

	queue := make([]*BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmpSum := 0
		tmpLen := 0
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			tmpSum += node.Data
			tmpLen += 1

			if node.LeftNode != nil {
				queue = append(queue, node.LeftNode)
			}

			if node.RightNode != nil {
				queue = append(queue, node.RightNode)
			}

		}

		res = append(res, float64(tmpSum) / float64(tmpLen))
	} 

	return res
}