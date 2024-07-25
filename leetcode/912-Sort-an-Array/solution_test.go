package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Example 2",
			nums:     []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sortArray(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
