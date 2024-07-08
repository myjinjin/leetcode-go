package leetcode

func findTheWinner(n int, k int) int {
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i + 1
	}

	for len(queue) > 1 {
		try := k
		for try > 1 {
			queue = append(queue[1:], queue[0])
			try--
		}
		queue = queue[1:]
	}
	return queue[0]
}
