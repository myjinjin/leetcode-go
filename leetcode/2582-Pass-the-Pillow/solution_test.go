package leetcode

import "testing"

func TestPassThePillow(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		time     int
		expected int
	}{
		{
			name:     "Example 1",
			n:        4,
			time:     5,
			expected: 2,
		},
		{
			name:     "Example 2",
			n:        3,
			time:     2,
			expected: 3,
		},
		{
			name:     "Example 3",
			n:        8,
			time:     9,
			expected: 6,
		},
		{
			name:     "Example 4",
			n:        18,
			time:     38,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := passThePillow(tc.n, tc.time)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
