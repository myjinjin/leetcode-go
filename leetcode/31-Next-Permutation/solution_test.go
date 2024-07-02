package leetcode

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextPermutation(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, tc.nums)
			}
		})
	}
}
