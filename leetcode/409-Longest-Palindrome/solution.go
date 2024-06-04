package leetcode

func longestPalindrome(s string) int {
	counts := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	hasOdd := false
	palindrome := 0
	for _, v := range counts {
		if v%2 == 0 {
			palindrome += v
		} else {
			palindrome += v - 1
			hasOdd = true
		}
	}
	if hasOdd {
		palindrome++
	}

	return palindrome
}
