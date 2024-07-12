package leetcode

func maximumGain(s string, x int, y int) int {
	removeCombo := func(str string, a, b byte, score int) (int, string) {
		stack := []byte{}
		gain := 0
		for i := 0; i < len(str); i++ {
			if len(stack) > 0 && stack[len(stack)-1] == a && str[i] == b {
				stack = stack[:len(stack)-1]
				gain += score
			} else {
				stack = append(stack, str[i])
			}
		}
		return gain, string(stack)
	}

	var totalGain int
	if x > y {
		gain1, remain := removeCombo(s, 'a', 'b', x)
		gain1Extra, _ := removeCombo(remain, 'b', 'a', y)
		totalGain = gain1 + gain1Extra
	} else {
		gain2, remain2 := removeCombo(s, 'b', 'a', y)
		gain2Extra, _ := removeCombo(remain2, 'a', 'b', x)
		totalGain = gain2 + gain2Extra
	}
	return totalGain
}
