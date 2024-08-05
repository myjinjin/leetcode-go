package leetcode

func minimumCost(s string) int64 {
	var result int64
	n := len(s)

	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			if i+1 <= n-i-1 {
				result += int64(i + 1)
			} else {
				result += int64(n - i - 1)
			}
		}
	}

	return result
}
