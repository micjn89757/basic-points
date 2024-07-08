package linkedlist


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头节点
    dumpHead := &ListNode{}
    dumpHead.Next = head
    pre := dumpHead
    cur := dumpHead

    // 移动cur节点
    for i := 0; i < n; i++ {
        if cur.Next != nil {
            cur = cur.Next
        }
    }

    for cur.Next != nil {
        cur = cur.Next
        pre = pre.Next
    }

    // 删除节点
    pre.Next = pre.Next.Next
    return dumpHead.Next
}