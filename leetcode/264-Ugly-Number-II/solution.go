package leetcode

func nthUglyNumber(n int) int {
	uglyNumber := make([]int, n)
	uglyNumber[0] = 1

	twoIndex, threeIndex, fiveIndex := 0, 0, 0
	two, three, five := 2, 3, 5

	for i := 1; i < n; i++ {
		uglyNumber[i] = min(two, min(three, five))

		if uglyNumber[i] == two {
			twoIndex++
			two = uglyNumber[twoIndex] * 2
		}
		if uglyNumber[i] == three {
			threeIndex++
			three = uglyNumber[threeIndex] * 3
		}
		if uglyNumber[i] == five {
			fiveIndex++
			five = uglyNumber[fiveIndex] * 5
		}
	}

	return uglyNumber[n-1]
}
