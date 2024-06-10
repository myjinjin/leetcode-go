package leetcode

import "testing"

func TestHeightChecker(t *testing.T) {
	testCases := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
		{
			name:     "Example 2",
			heights:  []int{5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Example 3",
			heights:  []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := heightChecker(tc.heights)
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}
