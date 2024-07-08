package leetcode

func passThePillow(n int, time int) int {
	cycle := 2*n - 2
	time %= cycle

	if time < n {
		return time + 1
	}

	return 2*n - time - 1
}
