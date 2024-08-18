package leetcode

import "testing"

func TestCountCollisions(t *testing.T) {
	testCases := []struct {
		name       string
		directions string
		expected   int
	}{
		{
			name:       "Example 1",
			directions: "RLRSLL",
			expected:   5,
		},
		{
			name:       "Example 2",
			directions: "LLRR",
			expected:   0,
		},
		{
			name:       "Example 3",
			directions: "RSLLRSSL",
			expected:   5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countCollisions(tc.directions)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
