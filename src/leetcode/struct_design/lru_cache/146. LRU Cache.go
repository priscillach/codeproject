package lru_cache

type Node struct {
	prev, next *Node
	Key, Value int
}

type LinkedList struct {
	head, tail *Node
	Len        int
}

func initLinkedList() *LinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LinkedList{
		head: head,
		tail: tail,
	}
}

func (l *LinkedList) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil
	l.Len--
}

func (l *LinkedList) pushFront(node *Node) {
	next := l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = next
	next.prev = node
	l.Len++
}

type LRUCache struct {
	linkedList *LinkedList
	cacheMap   map[int]*Node
	cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		linkedList: initLinkedList(),
		cap:        capacity,
		cacheMap:   make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		this.linkedList.remove(node)
		this.linkedList.pushFront(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cacheMap[key]; ok {
		node.Value = value
		this.linkedList.remove(node)
		this.linkedList.pushFront(node)
	} else {
		node = &Node{
			Key:   key,
			Value: value,
		}
		this.linkedList.pushFront(node)
		this.cacheMap[key] = node

	}
	if this.linkedList.Len > this.cap {
		last := this.linkedList.tail.prev
		this.linkedList.remove(last)
		delete(this.cacheMap, last.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
