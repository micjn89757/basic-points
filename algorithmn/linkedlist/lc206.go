package linkedlist

type Lc206ListNode struct {
	Val int 
	Next *Lc206ListNode
}


func reverseList(head *Lc206ListNode) *Lc206ListNode {
	cur := head 
	var pre *Lc206ListNode

	for head != nil {
		head = head.Next
		cur.Next = pre 
		pre = cur 
		cur = head
	}

	return pre
}