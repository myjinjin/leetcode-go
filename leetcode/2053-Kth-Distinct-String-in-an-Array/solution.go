package leetcode

func kthDistinct(arr []string, k int) string {
	existMap := make(map[string]int)
	for _, s := range arr {
		existMap[s]++
	}

	for i := 0; i < len(arr); i++ {
		if existMap[arr[i]] == 1 {
			k--
			if k == 0 {
				return arr[i]
			}
		}
	}

	return ""
}
