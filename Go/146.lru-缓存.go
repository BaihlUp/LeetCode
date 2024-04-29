/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	cap        int
	cache      map[int]*DLinkNode
	size       int
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, value int
	pre, next  *DLinkNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkNode{},
		head:  &DLinkNode{key: 0, value: 0},
		tail:  &DLinkNode{key: 0, value: 0},
		cap:   capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.MoveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		if this.size+1 > this.cap {
			removeNode := this.DeleteTail()
			delete(this.cache, removeNode.key)
			this.size--
		}
		node := &DLinkNode{key: key, value: value}
		this.cache[key] = node
		this.AddToHead(node)
		this.size++
	} else {
		node.value = value
		this.MoveToHead(node)
	}
}

func (this *LRUCache) AddToHead(node *DLinkNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) MoveToHead(node *DLinkNode) {
	this.DeleteNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) DeleteNode(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) DeleteTail() *DLinkNode {
	node := this.tail.pre
	this.DeleteNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

