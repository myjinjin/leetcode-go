package leetcode

func divisibilityArray(word string, m int) []int {
	n := len(word)
	divisibility := make([]int, n)

	num := 0
	for i := 0; i < n; i++ {
		num *= 10
		num += int(word[i] - '0')
		num %= m

		if num == 0 {
			divisibility[i] = 1
		}
	}

	return divisibility
}
