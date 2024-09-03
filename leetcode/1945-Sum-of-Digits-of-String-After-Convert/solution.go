package leetcode

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	numStr := ""
	for i := 0; i < len(s); i++ {
		n := int(s[i] - 'a' + 1)
		numStr += strconv.Itoa(n)
	}

	result := 0
	for k > 0 {
		k--
		newNum := 0
		for i := 0; i < len(numStr); i++ {
			n, _ := strconv.Atoi(string(numStr[i]))
			newNum += n
		}
		result = newNum
		numStr = fmt.Sprintf("%d", newNum)
	}

	return result
}
