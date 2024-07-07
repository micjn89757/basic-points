package linkedlist

// 删除链表中等于给定值 val 的所有节点。

type Lc209ListNode struct {
	Val  int
	Next *Lc209ListNode
}

func removeElements(head *Lc209ListNode, val int) *Lc209ListNode {
	p := head

	// 处理头节点
	for head != nil && head.Val == val {
		head = head.Next
	}

	// 处理非头节点
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			// 注意删除之后不要直接移动，因为下一个元素还没有判断
			p = p.Next
		}
	}

	return p
}
