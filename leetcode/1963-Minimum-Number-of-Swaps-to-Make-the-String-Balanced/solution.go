package leetcode

func minSwaps(s string) int {
	open := 0
	unbalanced := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			open++
		} else if s[i] == ']' {
			if open > 0 {
				open--
			} else {
				unbalanced++
			}
		}
	}
	return (unbalanced + 1) / 2
}
