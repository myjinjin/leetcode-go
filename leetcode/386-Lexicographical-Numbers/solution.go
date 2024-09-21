package leetcode

func lexicalOrder(n int) []int {
	result := make([]int, n)
	curr := 1
	for i := 0; i < n; i++ {
		result[i] = curr
		if curr*10 <= n {
			curr *= 10
		} else {
			for curr%10 == 9 || curr+1 > n {
				curr /= 10
			}
			curr++
		}
	}
	return result
}
