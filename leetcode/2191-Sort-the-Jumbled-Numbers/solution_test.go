package leetcode

import (
	"reflect"
	"testing"
)

func TestSortJumbled(t *testing.T) {
	testCases := []struct {
		name     string
		mapping  []int
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			mapping:  []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6},
			nums:     []int{991, 338, 38},
			expected: []int{338, 38, 991},
		},
		{
			name:     "Example 2",
			mapping:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			nums:     []int{789, 456, 123},
			expected: []int{123, 456, 789},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sortJumbled(tc.mapping, tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
