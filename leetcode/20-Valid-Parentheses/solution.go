package leetcode

func isValid(s string) bool {
	stack := []rune{}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) > 0 && stack[len(stack)-1] == getOpenPair(r) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func getOpenPair(closeBracket rune) rune {
	switch closeBracket {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	}
	return 0
}
