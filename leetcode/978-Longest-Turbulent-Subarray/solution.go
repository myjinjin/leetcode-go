package leetcode

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	maxSize := 1
	up := 1
	down := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			up = down + 1
			down = 1
		} else if arr[i] < arr[i-1] {
			down = up + 1
			up = 1
		} else {
			up = 1
			down = 1
		}
		maxSize = max(maxSize, max(up, down))
	}

	return maxSize
}
