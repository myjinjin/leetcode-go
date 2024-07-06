package leetcode

import "testing"

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "Example 1",
			x:        4,
			expected: 2,
		},
		{
			name:     "Example 2",
			x:        8,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mySqrt(tc.x)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
