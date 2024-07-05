package leetcode

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "Example 1",
			num:      16,
			expected: true,
		},
		{
			name:     "Example 2",
			num:      14,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPerfectSquare(tc.num)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
