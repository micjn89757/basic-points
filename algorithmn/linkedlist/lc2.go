package linkedlist


// 两数相加 medium
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val int 
	Next *ListNode
}

// 注意：循环条件，进位处理，指针移动状态。
// addTwoNumbers 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addNum := 0 
	p1 := l1
	p2 := l2
	n := 0 
	resList := &ListNode{}
	p3 := resList

	// 注意循环条件
	for p1 != nil && p2 != nil {
		// 本位求和
		tempSum := p1.Val + p2.Val + addNum
		tempNum := tempSum % 10
		addNum = tempSum / 10

		// 保存结果
		if n == 0 {
			p3.Val = tempNum
			p3.Next = nil
		}else {
			node := &ListNode{
				Val: tempNum,
				Next: nil,
			}

			p3.Next = node
			p3 = p3.Next
		}

		n++
		// 移动指针
		p1 = p1.Next
		p2 = p2.Next
	}


	// 若l1或l2仍然有剩余数字, 此时p3在resList最后一个node
	// 注意进位和循环条件
	for p1 != nil {
		tempSum := addNum + p1.Val
		tempNum := tempSum % 10
		addNum = tempSum / 10
		node := &ListNode{
			Val: tempNum,
			Next: nil,
		}
		p3.Next = node 
		p3 = p3.Next
		p1 = p1.Next
	}

	for p2 != nil {
		tempSum := addNum + p2.Val
		tempNum := tempSum % 10
		addNum = tempSum / 10

		node := &ListNode{
			Val: tempNum,
			Next: nil,
		}

		p3.Next = node 
		p3 = p3.Next
		p2 = p2.Next
	}
	
	// 进位不为空
	if addNum != 0 {
		node := &ListNode{
			Val: addNum,
			Next: nil,
		}
		p3.Next = node
	}

	return resList
}
