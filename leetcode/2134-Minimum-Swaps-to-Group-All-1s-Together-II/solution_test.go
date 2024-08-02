package leetcode

import (
	"reflect"
	"testing"
)

func TestMinSwaps(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 1, 1, 0, 0},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 0, 0, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minSwaps(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
