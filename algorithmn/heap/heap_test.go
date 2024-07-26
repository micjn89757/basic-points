package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewIntHeap(MaxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(1)
	t.Logf("%#v", heap)
	heap.Push(2)
	t.Logf("%#v", heap)
	heap.Push(3)
	t.Logf("%#v", heap)
	heap.Push(4)
	t.Logf("%#v", heap)
	heap.Push(5)
	t.Logf("%#v", heap)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	t.Log(heap.Pop()) 
	t.Log(heap.Pop()) 
	t.Log(heap.Pop()) 
	t.Log(heap.Pop())        
	t.Log(heap.Pop())        

	/* 获取堆大小 */
	t.Logf("堆元素数量为 %d\n", heap.Size())
}
