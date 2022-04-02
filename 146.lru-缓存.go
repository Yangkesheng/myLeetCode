/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

import "container/list"

// @lc code=start
type LRUCache struct {
	data     map[int]*list.Element
	capacity int
	hits     *list.List
}

type kv struct {
	k, v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		data:     make(map[int]*list.Element),
		capacity: capacity,
		hits:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if el, ok := this.data[key]; ok {
		this.hits.MoveToBack(el)
		return el.Value.(kv).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if el, ok := this.data[key]; ok {
		this.hits.MoveToBack(el)
		el.Value = kv{key, value}
		return
	}

	if this.hits.Len() == this.capacity {
		delete(this.data, this.hits.Front().Value.(kv).k)
		this.hits.Remove(this.hits.Front())
	}

	this.data[key] = this.hits.PushBack(kv{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
