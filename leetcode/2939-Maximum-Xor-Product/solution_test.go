package leetcode

import "testing"

func TestMaximumXorProduct(t *testing.T) {
	testCases := []struct {
		name     string
		a        int64
		b        int64
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			a:        12,
			b:        5,
			n:        4,
			expected: 98,
		},
		{
			name:     "Example 2",
			a:        6,
			b:        7,
			n:        5,
			expected: 930,
		},
		{
			name:     "Example 3",
			a:        1,
			b:        6,
			n:        3,
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumXorProduct(tc.a, tc.b, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
