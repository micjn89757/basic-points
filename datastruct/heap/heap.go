/*
堆是一种完全二叉树（最底层节点靠左填充，其他层的节点都被填满）
大根堆：任意节点的值<=其子节点的值，根节点值是最大的
小根堆：任意节点的值>=其子节点的值，根节点值是最大的

常用于第k大，第k小的排序任务

堆常常用于实现编程语言中提供的优先队列

堆排序的最坏、最好、平均时间复杂度为O(nlogn)


Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆，实现 heap.Interface 需要同时实现 sort.Interface
*/
package heap

type Heap interface {
	Push(any) 			// 元素入堆 时间复杂度:O(logn)
	Pop() 	any 	// 堆顶元素出堆 时间复杂度: O(logn)
	Peek() 	any		// 访问堆顶元素 O(1)
	Size()	int 	// 获取堆的元素数量 O(1)
	Empty() bool	// 判断堆是否为空 	O(1)
}

type HeapType uint

const (
	MinHeap HeapType = iota 
	MaxHeap
)

type IntHeap struct {
	Flag 	HeapType	// 控制大根堆小根堆
	Data	[]int 
}

func NewIntHeap(flag HeapType) *IntHeap {
	return &IntHeap{
		Flag: flag,
		Data: []int{},
	}
}

// !Push 元素入堆, 先将元素插入堆底(最后一个元素)，向上逐渐堆顶化（与父节点对比）
func (h *IntHeap) Push(el int) {
	h.Data = append(h.Data, el)

	elIndex := len(h.Data) - 1 // 插入位置
	for (elIndex - 1) / 2 >= 0 {	// 堆化最多调整至根节点
		pIndex := (elIndex - 1) / 2

		if h.Flag == MinHeap {
			if h.Data[pIndex] > h.Data[elIndex] {
				h.Swap(pIndex, elIndex)
			} else {
				return
			}
		} else {
			if h.Data[pIndex] < h.Data[elIndex] {
				h.Swap(pIndex, elIndex)
			} else {
				return
			}
		}

		elIndex = pIndex
	}

}

func (h *IntHeap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
} 

// !Pop 出堆，交换堆顶与堆底元素，在再将堆底元素删除，再进行堆底化（自顶向下）
func (h *IntHeap) Pop() int {
	// 元素出堆，交换堆顶堆底，删除堆底
	val := h.Data[0]
	h.Swap(0, h.Size() - 1)
	h.Data = h.Data[:h.Size() - 1]

	// 堆底化
	index := 0
	mmIndex := index 
	if h.Flag == MinHeap {
		for {
			left  := index * 2 + 1
			right := index * 2 + 2
			if left < h.Size() && h.Data[left] < h.Data[index] {
				mmIndex = left
			} 
			
			if right < h.Size() && h.Data[right] < h.Data[left] {
				mmIndex = right
			} 
			
			if mmIndex == index {
				break
			}

			// 交换节点
			h.Swap(mmIndex, index)

			// 向下堆化
			index = mmIndex

		}
	} else {
		for {
			left  := index * 2 + 1
			right := index * 2 + 2
			if left < h.Size() && h.Data[left] > h.Data[index] {
				mmIndex = left
			} 
			
			if right < h.Size() && h.Data[right] > h.Data[left] {
				mmIndex = right
			} 
			
			if mmIndex == index {
				break
			}

			// 交换节点
			h.Swap(mmIndex, index)

			// 向下堆化
			index = mmIndex

		}
	}


	return val
}

func (h *IntHeap) Peek() int {
	return h.Data[0]
}

func (h *IntHeap) Size() int {
	return len(h.Data)
}
