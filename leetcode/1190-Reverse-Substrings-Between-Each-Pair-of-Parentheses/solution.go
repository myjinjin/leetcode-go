package leetcode

func reverseParentheses(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			temp := make([]byte, 0)
			for len(stack) > 0 {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if popped == '(' {
					break
				}
				temp = append(temp, popped)
			}
			stack = append(stack, temp...)
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
