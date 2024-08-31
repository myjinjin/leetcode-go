package leetcode

const MOD = 1000000007

func maximumXorProduct(a int64, b int64, n int) int {
	for i := n - 1; i >= 0; i-- {
		mask := int64(1 << i)

		if a&mask != 0 && b&mask != 0 {
			continue
		} else if a&mask != 0 {
			if a > b {
				a ^= mask
				b |= mask
			}
		} else if b&mask != 0 {
			if a < b {
				a |= mask
				b ^= mask
			}
		} else {
			a |= mask
			b |= mask
		}
	}

	a %= MOD
	b %= MOD

	return int((a * b) % MOD)
}
