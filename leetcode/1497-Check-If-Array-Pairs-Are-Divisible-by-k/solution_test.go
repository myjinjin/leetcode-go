package leetcode

import "testing"

func TestCanArrange(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		k        int
		expected bool
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
			k:        5,
			expected: true,
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 3, 4, 5, 6},
			k:        7,
			expected: true,
		},
		{
			name:     "Example 3",
			arr:      []int{1, 2, 3, 4, 5, 6},
			k:        10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canArrange(tc.arr, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
