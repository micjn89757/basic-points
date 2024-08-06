package tree

func pathSum(root *BinaryTreeNode, targetSum int) [][]int {
	res := make([][]int, 0)

	var traver func(*BinaryTreeNode, int, []int)
	traver = func(cur *BinaryTreeNode, count int, tmp []int) {
		if cur.LeftNode == nil && cur.RightNode == nil && count == targetSum {
			tmpRes := make([]int, len(tmp))	// 一定注意不要传slice header，要重新做新切片
			copy(tmpRes, tmp)
			res = append(res, tmpRes)
			return
		}

		if cur.LeftNode == nil && cur.RightNode == nil {
			return
		}

		if cur.LeftNode != nil {
			tmp = append(tmp, cur.LeftNode.Data)
			count += cur.LeftNode.Data
			traver(cur.LeftNode, count, tmp)
			count -= cur.LeftNode.Data
			tmp = tmp[:len(tmp)-1]
		}

		if cur.RightNode != nil {
			tmp = append(tmp, cur.RightNode.Data)
			count += cur.RightNode.Data
			traver(cur.RightNode, count, tmp)
			count -= cur.RightNode.Data
			tmp = tmp[:len(tmp)-1]
		}

		return
	}

	traver(root, root.Data, []int{root.Data})

	return res
}
