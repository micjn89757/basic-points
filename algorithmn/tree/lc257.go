package tree

import "strconv"

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。
*/


func binaryTreePaths(root *BinaryTreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	traversal(root, []int{}, res)

	return res
}


// 前序遍历
func traversal(cur *BinaryTreeNode, path []int, res []string) {
	// path加入本节点
	path = append(path, cur.Data)
	// 终止条件，到达叶子节点
	if cur.LeftNode == nil && cur.RightNode == nil {
		var tmpPath string
		for i := 0; i < len(path) - 1; i++ {
			tmpPath += strconv.Itoa(path[i])
			tmpPath += "->"
		}

		tmpPath += strconv.Itoa(path[len(path) - 1])
		res = append(res, tmpPath)	// 找到一条路径
		return 
	}


	if cur.LeftNode != nil {
		traversal(cur.LeftNode, path, res)
		path = path[:len(path) - 1] // 回溯，因为子节点已经处理完了
	}

	if cur.RightNode != nil {
		traversal(cur.RightNode, path, res)
		path = path[:len(path) - 1]	// 回溯
	}
}