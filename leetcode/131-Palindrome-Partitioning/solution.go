package leetcode

func partition(s string) [][]string {
	result := [][]string{}
	var backtrack func(start int, curr []string)
	backtrack = func(start int, curr []string) {
		if start == len(s) {
			result = append(result, append([]string{}, curr...))
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s[start : end+1]) {
				curr = append(curr, s[start:end+1])
				backtrack(end+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}
	backtrack(0, []string{})
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
