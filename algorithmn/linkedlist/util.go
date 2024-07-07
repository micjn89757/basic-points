package linkedlist

// GenerateLinkedList 根据传入的数组生成链表
// TODO: 适配更多类型
func GenerateLinkedList(arr []int) *Lc206ListNode {
	if len(arr) == 0 {
		return nil
	}

	var head, p *Lc206ListNode
	for _, v :=  range arr {
		NewNode := &Lc206ListNode{
			Val: v,
			Next: nil,
		}

		if head == nil {
			head = NewNode
			p = head
		} else {
			p.Next = NewNode
			p = p.Next
		}
	}

	return head
}