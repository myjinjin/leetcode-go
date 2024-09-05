package leetcode

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)

	result := make([]int, n)

	sum := (n + m) * mean
	for i := 0; i < m; i++ {
		sum -= rolls[i]
	}
	if sum <= 0 || sum < n || sum > 6*n {
		return []int{}
	}

	base := sum / n
	remainder := sum % n
	for i := 0; i < n; i++ {
		result[i] = base
		if remainder > 0 {
			result[i]++
			remainder--
		}
	}

	return result
}
