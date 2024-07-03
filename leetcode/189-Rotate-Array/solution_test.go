package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.nums, tc.k)
			if !reflect.DeepEqual(tc.nums, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, tc.nums)
			}
		})
	}
}
