package leetcode

func minEatingSpeed(piles []int, h int) int {
	kMin := 1
	kMax := max(piles)

	for kMin < kMax {
		mid := (kMin + kMax) / 2
		totalTime := 0
		for _, pile := range piles {
			totalTime += pile / mid
			if pile%mid > 0 {
				totalTime += 1
			}
		}
		if totalTime <= h {
			kMax = mid
		} else {
			kMin = mid + 1
		}
	}
	return kMin
}

func max(arr []int) int {
	result := 0
	for _, n := range arr {
		if n > result {
			result = n
		}
	}
	return result
}
