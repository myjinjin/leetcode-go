package leetcode

import "math"

func judgeSquareSum(c int) bool {
	max := int(math.Sqrt(float64(c)))
	for i := 0; i <= max; i++ {
		a := i
		b := int(math.Sqrt(float64(c - a*a)))
		if a*a+b*b == c {
			return true
		}
	}
	return false
}
