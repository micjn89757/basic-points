package linkedlist 



func swapPairs(head *ListNode) *ListNode {
	// 创建虚拟头节点
	dumpHead := &ListNode{}
	dumpHead.Next = head 

	tmp := dumpHead

	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next 
		tmp.Next = node2 
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}

	return dumpHead.Next
}