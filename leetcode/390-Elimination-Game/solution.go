package leetcode

func lastRemaining(n int) int {
	leftmost := 1
	interval := 1
	remainingCount := n
	fromLeft := true

	for remainingCount > 1 {
		if fromLeft || remainingCount%2 == 1 {
			leftmost += interval
		}

		interval *= 2
		remainingCount /= 2
		fromLeft = !fromLeft
	}

	return leftmost
}
