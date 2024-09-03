package leetcode

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	dp := make(map[int]int64)
	mod := int64(1000000007)
	sort.Ints(arr)

	for i, n := range arr {
		dp[n] = 1
		for j := 0; j < i; j++ {
			factor := arr[j]
			if n%factor == 0 {
				quotient := n / factor
				if _, ok := dp[quotient]; ok {
					dp[n] = (dp[n] + int64(dp[factor])*int64(dp[quotient])) % mod
				}
			}
		}
	}

	var count int64
	for _, c := range dp {
		count = (count + c) % mod
	}
	return int(count)
}
