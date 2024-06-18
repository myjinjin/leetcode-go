package leetcode

import "testing"

func TestMinFlips(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		c        int
		expected int
	}{
		{
			name:     "Example 1",
			a:        2,
			b:        6,
			c:        5,
			expected: 3,
		},
		{
			name:     "Example 2",
			a:        4,
			b:        2,
			c:        7,
			expected: 1,
		},
		{
			name:     "Example 3",
			a:        1,
			b:        2,
			c:        3,
			expected: 0,
		},
		{
			name:     "Example 4",
			a:        8, // 1000
			b:        3, // 0011
			c:        5, // 0101
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minFlips(tc.a, tc.b, tc.c)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
