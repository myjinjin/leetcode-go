package leetcode

import "strconv"

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	maxDigits := make([]int, n)
	maxDigits[n-1] = n - 1

	for i := n - 2; i >= 0; i-- {
		if s[i] > s[maxDigits[i+1]] {
			maxDigits[i] = i
		} else {
			maxDigits[i] = maxDigits[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if s[i] < s[maxDigits[i]] {
			s[i], s[maxDigits[i]] = s[maxDigits[i]], s[i]
			break
		}
	}

	result, _ := strconv.Atoi(string(s))
	return result
}
