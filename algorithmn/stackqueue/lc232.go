package stackqueue

/*
用栈实现队列
*/

type MyQueue struct {
    stack1 []int 
    stack2 []int 
}


func ConstructorMyQueue() MyQueue {
    return MyQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

// 推到队列末尾
func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}

// 队列开头移除元素并返回
func (this *MyQueue) Pop() int {
	// 判断stack2是否为空
	if len(this.stack2) == 0 {
		s1Len := len(this.stack1)

		// 移动到stack2中
		for i := s1Len - 1; i >= 0; i-- {	
			this.stack2 = append(this.stack2, this.stack1[i])
		}

		this.stack1 = []int{}
	}


	// 移除开头元素并返回
	s2Len := len(this.stack2)
	ret := this.stack2[s2Len - 1]
	this.stack2 = this.stack2[:s2Len - 1]
	return ret
}

// 返回队列开头元素
func (this *MyQueue) Peek() int {
	// 判断stack2是否为空
	if len(this.stack2) == 0 {
		s1Len := len(this.stack1)

		// 移动到stack2中
		for i := s1Len - 1; i >= 0; i-- {	
			this.stack2 = append(this.stack2, this.stack1[i])
		}

		this.stack1 = []int{}
	}


	s2Len := len(this.stack2)
	ret := this.stack2[s2Len - 1]
	return ret
}

// 队列是否为空
func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	}

	return false
}

