package leetcode

import (
	"reflect"
	"testing"
)

func TestApplyOperations(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 2, 1, 1, 0},
			expected: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1},
			expected: []int{1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := applyOperations(tc.nums)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
