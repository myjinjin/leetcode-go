package leetcode

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 1, 3, 2, 5},
			expected: []int{3, 5},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 0},
			expected: []int{-1, 0},
		},
		{
			name:     "Example 3",
			nums:     []int{0, 1},
			expected: []int{1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := singleNumber(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
