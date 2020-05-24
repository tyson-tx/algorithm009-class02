// 641. 设计循环双端队列

type MyCircularDeque struct {
	cache    []int
	capacity int
	length   int
	front    int
	rear     int
}

// 构造函数,双端队列的大小为k。
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cache:    make([]int, k),
		capacity: k,
		front:    1,
	}
}

// 将一个元素添加到双端队列头部。 如果操作成功返回 true。
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.capacity {
		return false
	}
	this.length++
	this.front--
	if this.front == -1 {
		this.front = this.capacity-1
	}
	this.cache[this.front] = value
	return true
}

// 将一个元素添加到双端队列尾部。如果操作成功返回 true。
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.capacity {
		return false
	}
	this.length++
	this.rear++
	if this.rear == this.capacity {
		this.rear = 0
	}
	this.cache[this.rear] = value
	return true
}

// 从双端队列头部删除一个元素。 如果操作成功返回 true。
func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}
	this.length--
	this.front++
	if this.front == this.capacity {
		this.front = 0
	}
	return true
}

// 从双端队列尾部删除一个元素。如果操作成功返回 true。
func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}
	this.length--
	this.rear--
	if this.rear == -1 {
		this.rear = this.capacity-1
	}
	return true
}

// 从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.cache[this.front]
}

// 获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	return this.cache[this.rear]
}

// 检查双端队列是否为空。
func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

// 检查双端队列是否满了。
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}
