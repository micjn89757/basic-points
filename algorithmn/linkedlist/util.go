package linkedlist

type ListNode struct {
	Val int
	Next *ListNode
}

// GenerateLinkedList 根据传入的数组生成链表
func GenerateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var head, p *ListNode
	for _, v :=  range arr {
		NewNode := &ListNode{
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

