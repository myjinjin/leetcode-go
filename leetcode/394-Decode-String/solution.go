package leetcode

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []string{}
	currNum := 0
	currStr := ""

	for _, c := range s {
		if c == '[' {
			stack = append(stack, currStr)
			stack = append(stack, strconv.Itoa(currNum))
			currStr = ""
			currNum = 0
		} else if c == ']' {
			num, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currStr = prevStr + strings.Repeat(currStr, num)
		} else if c >= '0' && c <= '9' {
			currNum = currNum*10 + int(c-'0')
		} else {
			currStr += string(c)
		}
	}

	return currStr
}
