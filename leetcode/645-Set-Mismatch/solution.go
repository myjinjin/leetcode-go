package leetcode

func findErrorNums(nums []int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	duplicated, missing := -1, -1
	for i := 1; i <= len(nums); i++ {
		if val, exist := count[i]; !exist {
			missing = i
		} else if val == 2 {
			duplicated = i
		}
	}
	return []int{duplicated, missing}
}
