package main

import "fmt"

func main() {
	cache := newLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)
	//["LRUCache", "put",   "put", "get","put", "get", "put","get","get","get"]
	//	[[2],       [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	fmt.Println("DONE")
}

type Node struct {
	next *Node
	prev *Node
	key  int
	val  int
}

func newNode(key, val int) *Node {
	node := &Node{
		next: nil,
		prev: nil,
		key:  key,
		val:  val,
	}
	return node
}

type LRUCache struct {
	mp       map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func newLRUCache(capacity int) LRUCache {
	cache := LRUCache{
		mp:       make(map[int]*Node),
		head:     newNode(0, 0),
		tail:     newNode(0, 0),
		capacity: capacity,
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.mp[key]
	if !ok {
		return -1
	}
	remove(node, this)
	insert(node, this)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.mp[key]
	if ok {
		remove(node, this)
	}
	if this.capacity == len(this.mp) {
		remove(this.tail.prev, this)
	}
	node = newNode(key, value)
	insert(node, this)
}

func insert(node *Node, c *LRUCache) {
	c.mp[node.key] = node
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}
func remove(node *Node, c *LRUCache) {
	delete(c.mp, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}
