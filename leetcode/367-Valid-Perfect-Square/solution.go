package leetcode

func isPerfectSquare(num int) bool {
	left := 1
	right := num/2 + 1

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
