package leetcode

import "fmt"

func fractionAddition(expression string) string {
	i := 0
	numerator, denominator := 0, 1

	for i < len(expression) {
		sign := 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}

		var num int
		for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
			num *= 10
			num += int(expression[i] - '0')
			i++
		}
		num *= sign
		i++

		var den int
		for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
			den *= 10
			den += int(expression[i] - '0')
			i++
		}

		numerator = num*denominator + numerator*den
		denominator *= den
	}

	g := gcd(numerator, denominator)
	numerator /= g
	denominator /= g
	if denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}

	return fmt.Sprintf("%d/%d", numerator, denominator)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
