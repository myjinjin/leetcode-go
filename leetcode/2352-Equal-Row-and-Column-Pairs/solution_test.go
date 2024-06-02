package leetcode

import (
	"fmt"
	"testing"
)

func TestEqualPairs(t *testing.T) {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid: [][]int{
				{3, 2, 1},
				{1, 7, 6},
				{2, 7, 7},
			},
			expected: 1,
		},
		{
			grid: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.grid), func(t *testing.T) {
			result := equalPairs(tc.grid)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
