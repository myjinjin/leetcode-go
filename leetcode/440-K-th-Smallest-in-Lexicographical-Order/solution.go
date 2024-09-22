package leetcode

func findKthNumber(n int, k int) int {
	curr := 1
	k--

	for k > 0 {
		count := countNumbersInRange(n, curr, curr+1)
		if count <= k {
			curr++
			k -= count
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

func countNumbersInRange(maxNum, start, nextStart int) int {
	count := 0
	for start <= maxNum {
		count += min(maxNum+1, nextStart) - start
		start *= 10
		nextStart *= 10
	}
	return count
}
