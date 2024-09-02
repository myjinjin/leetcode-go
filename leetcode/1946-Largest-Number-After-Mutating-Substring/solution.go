package leetcode

func maximumNumber(num string, change []int) string {
	result := []byte(num)
	started := false

	for i := 0; i < len(num); i++ {
		n := int(num[i] - '0')
		if change[n] > n {
			result[i] = byte(change[n] + '0')
			started = true
		} else if change[n] < n && started {
			break
		}
	}

	return string(result)
}
