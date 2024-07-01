package leetcode

import "testing"

func TestThreeConsecutiveOdds(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "Example 1",
			arr:      []int{2, 6, 4, 1},
			expected: false,
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := threeConsecutiveOdds(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
