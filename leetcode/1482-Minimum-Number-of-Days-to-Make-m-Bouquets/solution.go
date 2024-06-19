package leetcode

func minDays(bloomDay []int, m int, k int) int {
	flowers := m * k
	if len(bloomDay) < flowers {
		return -1
	}
	result := -1
	left, right := 1, 1000000000
	for left <= right {
		mid := (left + right) / 2
		length, bouquets := 0, 0
		for _, bloom := range bloomDay {
			if bloom <= mid {
				length++
				if length >= k {
					length = 0
					bouquets++
				}
			} else {
				length = 0
			}
		}
		if bouquets >= m {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
