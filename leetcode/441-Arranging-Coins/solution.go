package leetcode

func arrangeCoins(n int) int {
	i := 1
	for ; n >= 0; i++ {
		n -= i
	}
	return i - 2
}
