package leetcode

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: [][2]int{},
	}
}

func (s *StockSpanner) Next(price int) int {
	result := 1
	for len(s.stack) > 0 && s.stack[len(s.stack)-1][0] <= price {
		result += s.stack[len(s.stack)-1][1]
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{price, result})
	return result
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
