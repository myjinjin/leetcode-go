package leetcode

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int)
	for i, n := range nums {
		numsMap[n] = append(numsMap[n], i)
	}

	result := []int{}
	for i, n := range nums {
		if v, ok := numsMap[target-n]; ok {
			if target-n == n {
				if len(v) == 1 {
					continue
				}
				result = []int{i, numsMap[target-n][1]}
				break
			} else {
				result = []int{i, numsMap[target-n][0]}
				break
			}
		}
	}

	return result
}
