package leetcode

func minSteps(n int) int {
	step := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		step[i] = i
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				step[i] = min(step[i], step[j]+i/j)
				break
			}
		}
	}
	return step[n]
}
