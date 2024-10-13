package leetcode

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 1},
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findErrorNums(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
