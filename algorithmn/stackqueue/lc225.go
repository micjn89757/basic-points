/*
用队列实现栈
*/
package stackqueue


type MyStack struct {
    queue   []int
}


func ConstructorMyStack() MyStack {
	return MyStack{
		queue: []int{},
	}
}


func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
	qLen := len(this.queue)	// 最后一个元素位置是qLen - 1
	res := -1

	// 每次去除元素qLen - 1
	for qLen != 0 {
		val := this.queue[0]
		if qLen == 1 {
			res = val
			this.queue = this.queue[1:] 
			break
		}
		this.queue = this.queue[1:] 
		this.queue = append(this.queue, val)
		qLen--
	}

	return res
}


func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}


func (this *MyStack) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}

	return false
}