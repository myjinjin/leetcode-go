package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{0},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			moveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, tc.nums)
			}
		})
	}
}
