/*
*	LRU最近最少使用
* 	1. 当我们连续插入A,B,C...Z，此时内存已经填满了
* 	2. 当我们插入6，那么此时内存存放时间最久的A会被淘汰
*   3. 当我们从外部读取数据C的时候，C就会被移动到头部，这时候C就是最晚被淘汰的
*   
*   问题：如果应用一次读取了大量数据，但是这些数据只会被读取一次，那么这些数据就会留存在缓存中很长时间，造成缓存污染
*/

package cache

import "container/list"


type entry struct {	
	key, value, freq int	// freq用于LFU
}

type LRUCache struct {
	list  		*list.List				// 双向循环链表
	capacity	int						// 限定的容量

	storage map[int]*list.Element	// 存储的key
}

// Get 获取元素
// 获取失败返回nil
// 获取成功：1. 将当前节点置于头节点 2. 返回对应的value
func (c *LRUCache) Get(key int) any {
	node, ok := c.storage[key]

	if ok {
		c.list.MoveToFront(node)
		return node.Value.(entry).value
	}

	return nil 
}


func (c *LRUCache) Put(key int, value int) {
	if node := c.storage[key]; node != nil { // 如果节点存在
		node.Value = entry{key: key, value: value}	// 更新节点的值
		c.list.MoveToFront(node)	// 移动到最前端
		return 
	}

	// 节点不存在
	c.storage[key] = c.list.PushFront(entry{key: key, value: value}) // 插入节点到头部
	
	if c.list.Len() > c.capacity {	// 容量拉满，删除最后一个节点
		delete(c.storage, c.list.Remove(c.list.Back()).(entry).key)
	}
}


