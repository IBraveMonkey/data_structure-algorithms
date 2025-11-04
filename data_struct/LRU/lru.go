package main

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	Capacity int
	Data     map[int]*Node
	Head     *Node
	Tail     *Node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func Constructor(capacity int) LRUCache {
	data := make(map[int]*Node, capacity)
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Capacity: capacity,
		Data:     data,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.addToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		this.remove(node)
		this.addToHead(node)
		node.Value = value
		return
	}

	if len(this.Data) >= this.Capacity {
		node := this.Tail.Prev
		this.remove(node)
		delete(this.Data, node.Key)
	}

	node = &Node{
		Key:   key,
		Value: value,
	}
	this.addToHead(node)
	this.Data[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
