package linkedlist

// 你可以选择使用单链表或者双链表，设计并实现自己的链表。

// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。

// 实现 MyLinkedList 类：

// MyLinkedList() 初始化 MyLinkedList 对象。
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。

// TODO: 可以改成双向链表
type Lc707LinkedList struct { 
	Val int
	Next *Lc707LinkedList
}


// 这里有虚拟头节点
func Constructor() Lc707LinkedList {
    return Lc707LinkedList{
        Val: -999,
        Next: nil,
    }
}


func (l *Lc707LinkedList) Get(index int) int {
    p := l.Next

    for p != nil && index >= 0 {
        if index == 0 {
            return p.Val
        }
        index--
        p = p.Next 
    }

    return -1
}


func (l *Lc707LinkedList) AddAtHead(val int)  {
    p := l.Next
    NewNode := Constructor()
    NewNode.Val = val 
    NewNode.Next = p 
    l.Next = &NewNode
}


func (l *Lc707LinkedList) AddAtTail(val int)  {
    NewNode := Constructor()
    NewNode.Val = val

    p := l 
    for p.Next != nil {
        p = p.Next
    }

    p.Next = &NewNode
    
}


func (l *Lc707LinkedList) AddAtIndex(index int, val int)  {
    NewNode := Constructor()
    NewNode.Val = val
    // 插入到index节点之前, 如果index等于链表长度，则追加到末尾，如果比长度更大，则不会插入
    p := l.Next
    target := index - 1

    // 头节点
    if index == 0 {
        l.AddAtHead(val)
        return 
    }

    for p != nil  {
        if target == 0 {
            NewNode.Next = p.Next
            p.Next = &NewNode
            break
        }
        target--
        p = p.Next
    }

}


func (l *Lc707LinkedList) DeleteAtIndex(index int)  {
    // 只剩虚拟头节点
    if l.Next == nil {
        return 
    }

    p := l.Next

    // 删除头节点
    if index == 0 {
        l.Next = p.Next
        return 
    }

    target := index - 1

    for p.Next != nil {
        if target == 0 {
            p.Next = p.Next.Next 
            break
        } else {
            target--
            p = p.Next
        }
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */