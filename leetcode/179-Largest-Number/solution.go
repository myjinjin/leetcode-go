package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))

	allZero := true
	for i, num := range nums {
		if num != 0 {
			allZero = false
		}
		strs[i] = strconv.Itoa(num)
	}

	if allZero {
		return "0"
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	return strings.Join(strs, "")
}
