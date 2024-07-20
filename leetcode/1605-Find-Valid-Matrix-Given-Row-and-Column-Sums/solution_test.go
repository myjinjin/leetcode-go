package leetcode

import (
	"reflect"
	"testing"
)

func TestRestoreMatrix(t *testing.T) {
	testCases := []struct {
		name     string
		rowSum   []int
		colSum   []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			rowSum:   []int{3, 8},
			colSum:   []int{4, 7},
			expected: [][]int{{3, 0}, {1, 7}},
		},
		{
			name:     "Example 2",
			rowSum:   []int{5, 7, 10},
			colSum:   []int{8, 6, 8},
			expected: [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := restoreMatrix(tc.rowSum, tc.colSum)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
