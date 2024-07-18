package leetcode

import (
	"reflect"
	"testing"
)

func TestUnmarkedSumArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected []int64
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 2, 1, 2, 3, 1},
			queries:  [][]int{{1, 2}, {3, 3}, {4, 2}},
			expected: []int64{8, 3, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 4, 2, 3},
			queries:  [][]int{{0, 1}},
			expected: []int64{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := unmarkedSumArray(tc.nums, tc.queries)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
