package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildMatrix(t *testing.T) {
	testCases := []struct {
		name          string
		k             int
		rowConditions [][]int
		colConditions [][]int
		expected      [][]int
	}{
		{
			name:          "Example 1",
			k:             3,
			rowConditions: [][]int{{1, 2}, {3, 2}},
			colConditions: [][]int{{2, 1}, {3, 2}},
			expected:      [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}},
		},
		{
			name:          "Example 2",
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}},
			colConditions: [][]int{{2, 1}},
			expected:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := buildMatrix(tc.k, tc.rowConditions, tc.colConditions)
			if reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
