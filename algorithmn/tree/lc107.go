package tree

/*
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

思路: 层序遍历后反转数组
*/

func levelOrderBottom(root *BinaryTreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	queue := make([]*BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmp := make([]int, 0)
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node.Data)

			if node.LeftNode != nil {
				queue = append(queue, node.LeftNode)
			}

			if node.RightNode != nil {
				queue = append(queue, node.RightNode)
			}

		}

		res = append(res, tmp)
	} 

	reverse(res)
	return res
}


func reverse(res [][]int) {

	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
	}
}
