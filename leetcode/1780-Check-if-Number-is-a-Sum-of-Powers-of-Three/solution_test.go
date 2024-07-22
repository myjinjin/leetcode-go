package leetcode

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Example 1",
			n:        12,
			expected: true,
		},
		{
			name:     "Example 2",
			n:        91,
			expected: true,
		},
		{
			name:     "Example 3",
			n:        21,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkPowersOfThree(tc.n)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
