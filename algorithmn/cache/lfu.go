/*
* LFU 最近最不常使用
* 如果数据过去被访问多次，那么将来被访问的频率也更高
* 会记录每个数据的访问次数。当一个数据被再次访问时，就会增加该数据的访问次数。这样就解决了偶尔被访问一次之后，数据留存在缓存中很长一段时间的问题，相比于 LRU 算法也更合理一些
 */
package cache

import "container/list"

type LFUCache struct {
	freqToList 	map[int]*list.List		// key记录频次，value存放对应的双向链表，每个双向链表都有一个哨兵
	keyToNode  	map[int]*list.Element	// 记录缓存中的数据

	capacity	int						// 容量
	minFreq		int 					// 最小的访问次数，因为有可能最低的访问次数是2，移除的时候就要移除2的最后一个元素
}


func Init(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		keyToNode: map[int]*list.Element{},
		freqToList: map[int]*list.List{},
	}
}

// 插入数据以及移动操作
func (c *LFUCache) pushFront(e *entry) {
	// 查看value的频次是否存在
	if _, ok := c.freqToList[e.freq]; !ok {
		c.freqToList[e.freq] = list.New()	// 双向链表
	}

	// 将节点插入到对应频次的链表头部，并保存数据
	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
}

// 访问数据（不论是更新还是查找）
func (c *LFUCache) getEntry(key int) *entry {
	node := c.keyToNode[key]
	if node == nil {	// 数据不存在
		return nil 
	}

	// 取出数据
	e := node.Value.(*entry)
	lst := c.freqToList[e.freq]  // 找到对应的频次
	lst.Remove(node)	// 取出数据
	if lst.Len() == 0 {	// 取出数据后，长度为空
		delete(c.freqToList, e.freq) // 移除空链表
		if c.minFreq == e.freq {	// 取出数据后，如果这个节点是频次最小节点，那么最小频次++
			c.minFreq++
		}
	}

	e.freq++ 	// 节点频次++
	c.pushFront(e) // 插入到对应频次的链表中
	return e
}

// 读取数据
func (c *LFUCache) Get(key, value int) int {
	if e := c.getEntry(key); e != nil {	// 数据存在
		return e.value
	}

	return -1
}


// 更新数据
func (c *LFUCache) Put(key, value int) {
	// getEntry已经包含移动动作了
	if e := c.getEntry(key); e != nil {	// 数据存在
		e.value = value 
		return 
	}

	// 数据达到阈值
	if len(c.keyToNode) == c.capacity { // 总数据量达到阈值
		lst := c.freqToList[c.minFreq]	// 频次最小的链表，删除最后一个节点
		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key)

		if lst.Len() == 0 { // 空了
			delete(c.freqToList, c.minFreq)	// 移除空链表
		}
	}

	c.pushFront(&entry{key, value, 1})	// 新书放到看过1次的链表的上面
	c.minFreq = 1
}	