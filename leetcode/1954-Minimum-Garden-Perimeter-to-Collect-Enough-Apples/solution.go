package leetcode

func minimumPerimeter(neededApples int64) int64 {
	var n int64 = 1
	var apples int64 = 0
	for apples < neededApples {
		apples = 2 * n * (n + 1) * (2*n + 1)
		if apples >= neededApples {
			return 8 * n
		}
		n++
	}
	return 8 * n
}
