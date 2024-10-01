package leetcode

func canArrange(arr []int, k int) bool {
	remainderCount := make(map[int]int)
	for _, num := range arr {
		remainder := ((num % k) + k) % k
		remainderCount[remainder]++
	}

	for remainder, count := range remainderCount {
		if remainder == 0 {
			if count%2 != 0 {
				return false
			}
		} else if remainder*2 == k {
			if count%2 != 0 {
				return false
			}
		} else {
			if count != remainderCount[k-remainder] {
				return false
			}
		}
	}
	return true
}
