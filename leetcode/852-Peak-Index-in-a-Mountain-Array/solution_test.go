package leetcode

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{0, 1, 0},
			expected: 1,
		},
		{
			name:     "Example 2",
			arr:      []int{0, 2, 1, 0},
			expected: 1,
		},
		{
			name:     "Example 3",
			arr:      []int{0, 10, 5, 2},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := peakIndexInMountainArray(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
