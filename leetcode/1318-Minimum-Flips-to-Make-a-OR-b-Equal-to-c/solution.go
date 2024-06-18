package leetcode

func minFlips(a int, b int, c int) int {
	flips := 0
	for a != 0 || b != 0 || c != 0 {
		lastA := a & 1
		lastB := b & 1
		lastC := c & 1
		if lastC != (lastA | lastB) {
			if lastC == 0 {
				if lastA == 1 {
					flips++
				}
				if lastB == 1 {
					flips++
				}
			} else if lastC == 1 {
				flips++
			}
		}
		c >>= 1
		b >>= 1
		a >>= 1
	}
	return flips
}
