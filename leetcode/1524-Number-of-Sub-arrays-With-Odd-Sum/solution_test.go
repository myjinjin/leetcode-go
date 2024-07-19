package leetcode

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 3, 5},
			expected: 4,
		},
		{
			name:     "Example 2",
			arr:      []int{2, 4, 6},
			expected: 0,
		},
		{
			name:     "Example 3",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			expected: 16,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numOfSubarrays(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
