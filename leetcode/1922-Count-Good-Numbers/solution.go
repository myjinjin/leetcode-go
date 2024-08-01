package leetcode

const mod = 1000000007

func countGoodNumbers(n int64) int {
	return int((power(5, (n+1)/2) * power(4, n/2)) % mod)
}

func power(x, n int64) int64 {
	if n == 0 {
		return 1
	}

	half := power(x, n/2)
	result := (half * half) % mod
	if n%2 == 1 {
		result = (result * x) % mod
	}
	return result
}
