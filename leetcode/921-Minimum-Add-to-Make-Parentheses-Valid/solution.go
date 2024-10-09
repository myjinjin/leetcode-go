package leetcode

func minAddToMakeValid(s string) int {
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			right++
		} else if right > 0 {
			right--
		} else {
			left++
		}
	}
	return left + right
}
