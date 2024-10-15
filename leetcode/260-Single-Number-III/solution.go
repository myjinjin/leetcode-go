package leetcode

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	rightMost := 1
	for xor&rightMost == 0 {
		rightMost <<= 1
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&rightMost != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}
