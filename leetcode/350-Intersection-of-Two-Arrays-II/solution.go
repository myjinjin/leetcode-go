package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}

	result := []int{}
	for _, num := range nums2 {
		if _, ok := nums1Map[num]; ok {
			result = append(result, num)
			nums1Map[num]--
			if nums1Map[num] == 0 {
				delete(nums1Map, num)
			}
		}
	}

	return result
}
