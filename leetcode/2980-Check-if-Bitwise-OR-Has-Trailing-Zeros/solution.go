package leetcode

func hasTrailingZeros(nums []int) bool {
	count := 0

	for _, n := range nums {
		if n&1 == 0 {
			count++
			if count == 2 {
				return true
			}
		}
	}

	return false
}
