package leetcode

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sortColors(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, tc.nums)
			}
		})
	}
}
