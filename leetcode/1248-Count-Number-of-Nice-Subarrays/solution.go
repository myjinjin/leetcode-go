package leetcode

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	result := 0
	oddCountFirstIndexMap := make(map[int]int)
	oddCountFirstIndexMap[0] = -1
	oddCount := 0
	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 {
			oddCount++
		}
		if _, ok := oddCountFirstIndexMap[oddCount]; !ok {
			oddCountFirstIndexMap[oddCount] = i
		}

		if oddCount >= k {
			result += oddCountFirstIndexMap[oddCount-k+1] - oddCountFirstIndexMap[oddCount-k]
		}
	}

	return result
}
