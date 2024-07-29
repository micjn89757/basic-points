package tree

import "math"

/*
层序遍历，找到每一行的最大值
*/

func largestValues(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	queue := make([]*BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmp := math.MinInt
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Data > tmp {
				tmp = node.Data
			}

			if node.LeftNode != nil {
				queue = append(queue, node.LeftNode)
			}

			if node.RightNode != nil {
				queue = append(queue, node.RightNode)
			}

		}

		res = append(res, tmp)
	} 

	return res
}