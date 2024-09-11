package leetcode

func minBitFlips(start int, goal int) int {
	result := 0
	for start > 0 || goal > 0 {
		if start&1 != goal&1 {
			result++
		}
		start >>= 1
		goal >>= 1
	}
	return result
}
