package leetcode

import (
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	return compute(expression, make(map[string][]int))
}

func compute(expression string, memo map[string][]int) []int {
	if result, ok := memo[expression]; ok {
		return result
	}

	var result []int
	for i, char := range expression {
		if char == '+' || char == '-' || char == '*' {
			left := compute(expression[:i], memo)
			right := compute(expression[i+1:], memo)

			for _, l := range left {
				for _, r := range right {
					switch char {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
	}

	memo[expression] = result
	return result
}
