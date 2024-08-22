package leetcode

func findComplement(num int) int {
	tmp := num
	i := 0
	for tmp > 0 {
		tmp >>= 1
		i++
	}

	mask := 1
	for i > 0 {
		mask <<= 1
		i--
	}
	mask--
	num ^= mask
	return num
}
