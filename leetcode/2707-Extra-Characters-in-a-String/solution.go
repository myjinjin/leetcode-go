package leetcode

func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool)
	for _, word := range dictionary {
		dict[word] = true
	}

	// dp[i]: s의 처음부터 i번째 문자까지 최소 필요한 extra 문자 수
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i < n+1; i++ {
		dp[i] = dp[i-1] + 1

		for j := 0; j < i; j++ {
			if dict[s[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}
