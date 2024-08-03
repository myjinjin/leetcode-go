package leetcode

import "reflect"

func canBeEqual(target []int, arr []int) bool {
	targetMap := make(map[int]int)
	arrMap := make(map[int]int)

	for _, n := range target {
		targetMap[n]++
	}
	for _, n := range arr {
		arrMap[n]++
	}
	return reflect.DeepEqual(targetMap, arrMap)
}
