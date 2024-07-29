package tree


/*
N叉树层序遍历
*/

type Node struct {
	Val int 
	Children []*Node
}


func NLevelOrder(root *Node) [][]int {
    if root == nil {
		return nil
	}
	res := make([][]int, 0)

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmp := make([]int, 0)
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node.Val)

			for _, v := range node.Children {
                if v != nil {
                    queue = append(queue, v)
                }
            }

		}

		res = append(res, tmp)
	} 

	return res
}
