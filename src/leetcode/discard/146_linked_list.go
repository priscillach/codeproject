package leetcode

import "container/list"

type LRUCache struct {
	usedQueue *list.List
	cacheMap  map[int]*list.Element
	capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		usedQueue: list.New(),
		cacheMap:  make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cacheMap[key]; ok {
		this.usedQueue.Remove(v)
		this.usedQueue.PushFront(v.Value)
		return v.Value.([]int)[1]
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cacheMap[key]; ok {
		this.usedQueue.Remove(v)
		this.usedQueue.PushFront([]int{key, value})
		this.cacheMap[key] = this.usedQueue.Front()
	} else {
		this.usedQueue.PushFront([]int{key, value})
		this.cacheMap[key] = this.usedQueue.Front()

	}
	if this.usedQueue.Len() > this.capacity {
		toDelete := this.usedQueue.Back()
		this.usedQueue.Remove(toDelete)
		delete(this.cacheMap, toDelete.Value.([]int)[0])
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
