package leetcode

func maxDistance(arrays [][]int) int {
	d := 0

	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		arrMin := arr[0]
		arrMax := arr[len(arr)-1]
		d = max(d, abs(maxVal-arrMin))
		d = max(d, abs(arrMax-minVal))
		minVal = min(minVal, arrMin)
		maxVal = max(maxVal, arrMax)
	}
	return d
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
