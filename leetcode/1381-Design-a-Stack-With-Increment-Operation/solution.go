package leetcode

type CustomStack struct {
	stack []int
	cap   int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0, maxSize),
		cap:   maxSize,
	}
}

func (s *CustomStack) Push(x int) {
	if len(s.stack) < s.cap {
		s.stack = append(s.stack, x)
	}
}

func (s *CustomStack) Pop() int {
	if len(s.stack) == 0 {
		return -1
	}
	result := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return result
}

func (s *CustomStack) Increment(k int, val int) {
	l := min(k, len(s.stack))
	for i := 0; i < l; i++ {
		s.stack[i] += val
	}
}
