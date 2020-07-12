package Week_01

//leetcode 641
//设计双端循环队列(链表解法)
type MyCircularDeque struct {
	head *Node
	tail *Node
	c    int //容量
	size int //当前长度
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := &Node{
		prev: nil,
		next: nil,
	}
	tail := &Node{
		prev: nil,
		next: nil,
	}
	head.next = tail
	head.prev = tail
	tail.prev = head
	tail.next = head
	return MyCircularDeque{
		head: head, //伪节点
		tail: tail, //伪节点
		c:    k,
		size: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{
		value: value,
		prev:  this.head,
		next:  this.head.next,
	}
	this.head.next.prev = node
	this.head.next = node
	this.size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{
		value: value,
		prev:  this.tail.prev,
		next:  this.tail,
	}
	this.tail.prev.next = node
	this.tail.prev = node
	this.size++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head.next.next.prev = this.head
	this.head.next = this.head.next.next
	this.size--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail.prev.prev.next = this.tail
	this.tail.prev = this.tail.prev.prev
	this.size--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.value
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.prev.value
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.size == this.c {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque2 object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
