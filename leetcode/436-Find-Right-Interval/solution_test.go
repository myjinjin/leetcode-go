package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRightInterval(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		expected  []int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 2}},
			expected:  []int{-1},
		},
		{
			name:      "Example 2",
			intervals: [][]int{{3, 4}, {2, 3}, {1, 2}},
			expected:  []int{-1, 0, 1},
		},
		{
			name:      "Example 3",
			intervals: [][]int{{1, 4}, {2, 3}, {3, 4}},
			expected:  []int{-1, 2, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findRightInterval(tc.intervals)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
