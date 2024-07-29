package tree

/*
这道题和lc117一样
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
*/

type Lc116Node struct {
	Val int
	Left *	Lc116Node
	Right *Lc116Node
 	Next *Lc116Node
}

func connect(root *Lc116Node) *Lc116Node {
	if root == nil {
		return nil
	}

	queue := make([]*Lc116Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

		tmp := make([]*Lc116Node, 0)
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		

		// 连接node
		for i := 0; i < len(tmp) - 1; i++ {
			tmp[i].Next = tmp[i + 1]
		}
	} 

	return root
}

// 法2
func connect2(root *Lc116Node) *Lc116Node {
	if root == nil {
		return nil
	}

	queue := make([]*Lc116Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)

        var nodePre *Lc116Node
        var node *Lc116Node
		// length表示当前层的节点个数
		for i := 0; i < length; i++ {		
            if i == 0 {
                node = queue[0]
			    queue = queue[1:]
                nodePre = node
            } else {
                node = queue[0]
			    queue = queue[1:]
                nodePre.Next = node 
                nodePre = node
            }


			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		
	} 

	return root
}