package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected int
	}{
		{
			name:     "Example 1",
			arr1:     []int{1, 10, 100},
			arr2:     []int{1000},
			expected: 3,
		},
		{
			name:     "Example 2",
			arr1:     []int{1, 2, 3},
			arr2:     []int{4, 4, 4},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestCommonPrefix(tc.arr1, tc.arr2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
