package linkedlist

func detectCycle(head *ListNode) *ListNode {
	// 双指针法， 首先找到相遇节点
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next 
		fast = fast.Next.Next

		// 快慢指针相遇
		if slow == fast {
			// 从起始点和相遇点相向而行
			index1 := head 
			index2 := fast

			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next.Next	
			}

			return index2
		}
	}

	return nil
}