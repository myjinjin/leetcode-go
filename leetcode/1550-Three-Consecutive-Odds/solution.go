package leetcode

func threeConsecutiveOdds(arr []int) bool {
	odds := 0
	for _, n := range arr {
		if n%2 == 1 {
			odds++
			if odds == 3 {
				return true
			}
		} else {
			odds = 0
		}
	}
	return false
}
