package redis

// quicklist实现

type quickListNode struct {
	prev *quickListNode
	next *quickListNode
	
}