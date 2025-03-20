/* package main

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	memory   map[int]*Node
	capacity int
	tail     *Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		memory:   make(map[int]*Node, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) moveToFront(node *Node) {
	if node == this.head {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == this.tail {
		this.tail = node.prev
	}

	node.next = this.head

	if this.head != nil {
		this.head.prev = node
	}

	this.head = node

	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.memory[key]; ok {
		this.moveToFront(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.memory[key]; ok {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.memory) >= this.capacity {

		lru := this.tail
		delete(this.memory, lru.key)
		this.tail = lru.prev

		if this.tail != nil {
			this.tail.next = nil
		} else {
			this.head = nil
		}
	}

	newNode := &Node{key: key, value: value}
	this.memory[key] = newNode
	this.moveToFront(newNode)
} */
