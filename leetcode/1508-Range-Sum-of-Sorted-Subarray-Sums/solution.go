package leetcode

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, n*(n+1)/2)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			sums = append(sums, sum)
		}
	}

	sort.Ints(sums)

	result := 0
	for i := left - 1; i <= right-1; i++ {
		result += sums[i]
		result %= 1000000007
	}
	return result
}
