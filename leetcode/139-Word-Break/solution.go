package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	for end := 1; end <= n; end++ {
		for start := 0; start < end; start++ {
			if dp[start] && wordSet[s[start:end]] {
				dp[end] = true
				break
			}
		}
	}
	return dp[n]
}
