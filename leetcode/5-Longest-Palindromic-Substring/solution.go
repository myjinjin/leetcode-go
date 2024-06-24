package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}
	}
	return s[start : start+maxLength]
}
