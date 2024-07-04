package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		curr := letters[mid]
		if curr > target {
			if mid-1 >= 0 && letters[mid-1] <= target {
				return curr
			}
			right = mid - 1
		} else if curr < target {
			left = mid + 1
		} else {
			if mid+1 < n && letters[mid+1] > target {
				return letters[mid+1]
			}
			left = mid + 1
		}
	}
	return letters[0]
}
