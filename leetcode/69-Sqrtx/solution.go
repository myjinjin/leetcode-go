package leetcode

func mySqrt(x int) int {
	left := 0
	right := x/2 + 1

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left*left > x {
		return left - 1
	}
	return left
}
