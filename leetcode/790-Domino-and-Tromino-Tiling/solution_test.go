package leetcode

import "testing"

func TestNumTilings(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        3,
			expected: 5,
		},
		{
			name:     "Example 2",
			n:        1,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numTilings(tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
