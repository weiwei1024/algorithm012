package Week_01

//leetcode 641
//使用静态数组实现双端循环队列(静态数组解法)
type MyCircularDeque2 struct {
	buffer []int
	k      int
	size   int
	front  int
	rear   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor2(k int) MyCircularDeque2 {
	return MyCircularDeque2{
		buffer: make([]int, k),
		size:   0,
		front:  k - 1,
		rear:   0,
		k:      k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque2) InsertFront(value int) bool {
	if this.size == this.k {
		return false
	}
	this.buffer[this.front] = value
	this.front = (this.front - 1 + this.k) % this.k //这里加上k为了保证非负
	this.size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque2) InsertLast(value int) bool {
	if this.size == this.k {
		return false
	}
	this.buffer[this.rear] = value
	this.rear = (this.rear + 1) % this.k
	this.size++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque2) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.k
	this.size--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque2) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.k) % this.k //这里加上k为了保证非负
	this.size--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque2) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[(this.front+1)%this.k] //front位置不是第一个元素
}

/** Get the last item from the deque. */
func (this *MyCircularDeque2) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[(this.rear-1+this.k)%this.k]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque2) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque2) IsFull() bool {
	if this.size == this.k {
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
