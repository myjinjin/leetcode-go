package leetcode

import (
	"reflect"
	"testing"
)

func TestLuckyNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			matrix:   [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			expected: []int{15},
		},
		{
			name:     "Example 2",
			matrix:   [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
			expected: []int{12},
		},
		{
			name:     "Example 3",
			matrix:   [][]int{{7, 8}, {1, 2}},
			expected: []int{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := luckyNumbers(tc.matrix)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
