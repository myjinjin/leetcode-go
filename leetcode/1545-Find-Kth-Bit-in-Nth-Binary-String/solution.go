package leetcode

import "math"

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	length := int(math.Pow(2.0, float64(n-1)) - 1)
	if k <= length {
		return findKthBit(n-1, k)
	}

	if k == length+1 {
		return '1'
	}

	if findKthBit(n-1, 2*(length+1)-k) == '0' {
		return '1'
	}
	return '0'
}
