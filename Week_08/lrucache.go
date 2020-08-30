package src

//双向链表+hash
type LRUCache struct {
	Capacity int
	Size     int                //当前
	Head     *CacheNode         //头节点
	Tail     *CacheNode         //尾节点
	Map      map[int]*CacheNode //记录节点
}

type CacheNode struct {
	Key  int
	Val  int
	Prev *CacheNode
	Next *CacheNode
}

func Constructor(capacity int) LRUCache {
	vHead := &CacheNode{Prev: nil} //虚拟节点
	vTail := &CacheNode{Next: nil} //虚拟节点
	vHead.Next = vTail
	vTail.Prev = vHead
	return LRUCache{
		Capacity: capacity,
		Size:     0,
		Head:     vHead,
		Tail:     vTail,
		Map:      make(map[int]*CacheNode, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Map[key]; ok {
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Prev = this.Head
		v.Next = this.Head.Next
		this.Head.Next.Prev = v
		this.Head.Next = v
		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Map[key]; ok {
		v.Val = value
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Prev = this.Head
		v.Next = this.Head.Next
		this.Head.Next.Prev = v
		this.Head.Next = v
		return
	}
	if this.Size == this.Capacity {
		delete(this.Map, this.Tail.Prev.Key)
		this.Tail.Prev.Prev.Next = this.Tail
		this.Tail.Prev = this.Tail.Prev.Prev
		this.Size--
	}
	n := &CacheNode{
		Key: key,
		Val: value,
	}
	n.Prev = this.Head
	n.Next = this.Head.Next
	this.Head.Next.Prev = n
	this.Head.Next = n
	this.Map[key] = n
	this.Size++
}
