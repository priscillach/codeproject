package lfu_cache

import "container/list"

type LFUCache struct {
	nodes        map[int]*list.Element
	lists        map[int]*list.List
	cap          int
	minFrequency int
}

type Node struct {
	Key int
	Val int
	Cnt int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes: make(map[int]*list.Element),
		lists: make(map[int]*list.List),
		cap:   capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.nodes[key]; !ok {
		return -1
	}
	element := this.nodes[key]
	node := element.Value.(*Node)
	// add cnt and minFrequency
	if node.Cnt == this.minFrequency && this.lists[node.Cnt].Len() == 1 {
		this.minFrequency++
	}
	this.lists[node.Cnt].Remove(element)
	if _, ok := this.lists[node.Cnt+1]; !ok {
		this.lists[node.Cnt+1] = list.New()
	}
	newElement := this.lists[node.Cnt+1].PushFront(node)
	this.nodes[key] = newElement
	node.Cnt++
	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.nodes[key]; ok {
		element := this.nodes[key]
		node := element.Value.(*Node)
		node.Val = value
		this.Get(key)
		return
	}

	if len(this.nodes) == this.cap {
		pop := this.lists[this.minFrequency].Back()
		delete(this.nodes, pop.Value.(*Node).Key)
		this.lists[this.minFrequency].Remove(pop)
	}

	node := &Node{
		Key: key,
		Val: value,
		Cnt: 1,
	}
	this.minFrequency = 1
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newElement := this.lists[1].PushFront(node)
	this.nodes[key] = newElement
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
