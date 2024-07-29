package leetcode

func numTeams(rating []int) int {
	count := 0
	n := len(rating)

	for j := 1; j < n-1; j++ {
		leftSmaller, leftLarger := 0, 0
		rightSmaller, rightLarger := 0, 0

		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				leftSmaller++
			} else if rating[i] > rating[j] {
				leftLarger++
			}
		}

		for k := j + 1; k < n; k++ {
			if rating[j] < rating[k] {
				rightLarger++
			} else if rating[j] > rating[k] {
				rightSmaller++
			}
		}

		count += leftSmaller*rightLarger + leftLarger*rightSmaller
	}

	return count
}
