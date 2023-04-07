package main

import "fmt"

func main() {
	n := 4
	inc(&n)
	print(n)
}
func inc(x *int) {
	*x++
}
func print(x int) {
	fmt.Println(x)
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
		remove(this.tail, this)
	}
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
