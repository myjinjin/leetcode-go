package leetcode

import "testing"

func TestCountGoodNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		n        int64
		expected int
	}{
		{
			name:     "Example 1",
			n:        1,
			expected: 5,
		},
		{
			name:     "Example 2",
			n:        4,
			expected: 400,
		},
		{
			name:     "Example 3",
			n:        50,
			expected: 564908303,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countGoodNumbers(tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
