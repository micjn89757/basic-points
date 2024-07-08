package linkedlist

// !给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 方法一使用map存储经过的路径
    // 双指针法: 既然是相交节点，相交节点的后续部分一定完全一致，毕竟相交节点对于两个列表来说就是同一个节点，那么也就是说相交节点以及后续部分一定是位于列表最后部分的。所以只需要调整长列表的起点，使得该起点到结尾的距离与短列表开头到结尾开头的距离一致，之后同步向后移，只要有相交节点那么必然能同时遇到。

    // 遍历得到两个链表的长度
    pA, pB := headA, headB 
    offset := 0 
    lenA, lenB := 0, 0
    for pA != nil {
        lenA++
        pA = pA.Next 
    }


    for pB != nil {
        lenB++
        pB = pB.Next 
    }

    pA = headA
    pB = headB

    if lenA > lenB {
        offset = lenA -lenB
        // 调整起始位置
        for offset != 0 {
            pA = pA.Next 
            offset--
        }

        for pA != nil {
            if pA == pB {
                return pA
            } 
            pA = pA.Next
            pB = pB.Next
        }

    } else {
        offset = lenB - lenA
        // 调整起始位置
        for offset != 0 {
            pB = pB.Next 
            offset--
        }

        for pB != nil {
            if pA == pB {
                return pB
            } 
            pA = pA.Next 
            pB = pB.Next
        }
    }


    return nil
}