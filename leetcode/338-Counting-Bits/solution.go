package leetcode

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		num := i
		for num != 0 {
			result[i] += num & 1
			num >>= 1
		}
	}
	return result
}
