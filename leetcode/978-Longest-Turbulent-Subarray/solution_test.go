package leetcode

import "testing"

func TestMaxTurbulenceSize(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			expected: 5,
		},
		{
			name:     "Example 2",
			arr:      []int{4, 8, 12, 16},
			expected: 2,
		},
		{
			name:     "Example 3",
			arr:      []int{100},
			expected: 1,
		},
		{
			name:     "Example 4",
			arr:      []int{9, 9},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxTurbulenceSize(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
