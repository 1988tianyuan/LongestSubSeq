package main

import (
	"container/list"
	"fmt"
)

//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
//
//
//
// 实现 LRUCache 类：
//
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
//限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//
//
//
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 3000
// 0 <= value <= 104
// 最多调用 3 * 104 次 get 和 put
//
// Related Topics 设计
// 👍 1074 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	queue *list.List
	cache map[int]*valueWrapper
	capacity int
	size int
}

type valueWrapper struct {
	value int
	node  *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		queue:    list.New(),
		cache:    make(map[int]*valueWrapper, capacity),
		capacity: capacity,
		size:     0,
	}
}


func (this *LRUCache) Get(key int) int {
	if this.cache[key] == nil {
		return -1
	}
	this.moveToHead(key)
	return this.cache[key].value
}


func (this *LRUCache) Put(key int, value int)  {
	if this.cache[key] != nil {
		this.cache[key].value = value
		this.moveToHead(key)
		return
	}
	if this.size == this.capacity {
		headE := this.queue.Front()
		delete(this.cache, headE.Value.(int))
		this.queue.Remove(headE)
		this.size--
	}
	e := this.queue.PushBack(key)
	this.cache[key] = &valueWrapper{value:value, node:e}
	this.size++
}

func (this *LRUCache) moveToHead(key int) {
	node := this.cache[key].node
	if node != nil {
		this.queue.Remove(node)
		newNode:= this.queue.PushBack(key)
		this.cache[key].node = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4,4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	cache2 := Constructor(2)
	fmt.Println(cache2.Get(2))
	cache2.Put(2, 6)
	fmt.Println(cache2.Get(1))
	cache2.Put(1, 5)
	cache2.Put(1,2)
	fmt.Println(cache2.Get(1))
	fmt.Println(cache2.Get(2))
}
