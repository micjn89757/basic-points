package linkedlist


func reverseList(head *ListNode) *ListNode {
	cur := head 
	var pre *ListNode

	for head != nil {
		head = head.Next
		cur.Next = pre 
		pre = cur 
		cur = head
	}

	return pre
}