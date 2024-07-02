package leetcode

type Node struct {
	value int
	min   int
	next  *Node
}

func NewNode(value, min int, next *Node) *Node {
	return &Node{
		value: value,
		min:   min,
		next:  next,
	}
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{head: nil}
}

func (s *MinStack) Push(val int) {
	if s.head == nil {
		s.head = NewNode(val, val, nil)
		return
	}

	min := s.head.min
	if min > val {
		min = val
	}
	s.head = NewNode(val, min, s.head)
}

func (s *MinStack) Pop() {
	s.head = s.head.next
}

func (s *MinStack) Top() int {
	return s.head.value
}

func (s *MinStack) GetMin() int {
	return s.head.min
}
