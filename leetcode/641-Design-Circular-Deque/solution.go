package leetcode

type MyCircularDeque struct {
	data     []int
	front    int
	rear     int
	size     int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:     make([]int, k),
		front:    0,
		rear:     k - 1,
		size:     0,
		capacity: k,
	}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.front = (q.front - 1 + q.capacity) % q.capacity
	q.data[q.front] = value
	q.size++
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % q.capacity
	q.size--
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear - 1 + q.capacity) % q.capacity
	q.size--
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.front]
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.rear]
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.size == q.capacity
}
