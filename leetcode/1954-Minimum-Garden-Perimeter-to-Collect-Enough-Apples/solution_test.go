package leetcode

import "testing"

func TestMinimumPerimeter(t *testing.T) {
	testCases := []struct {
		name         string
		neededApples int64
		expected     int64
	}{
		{
			name:         "Example 1",
			neededApples: 1,
			expected:     8,
		},
		{
			name:         "Example 2",
			neededApples: 13,
			expected:     16,
		},
		{
			name:         "Example 3",
			neededApples: 1000000000,
			expected:     5040,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumPerimeter(tc.neededApples)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
