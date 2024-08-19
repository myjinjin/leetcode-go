package leetcode

import "testing"

func TestMinSteps(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        3,
			expected: 3,
		},
		{
			name:     "Example 2",
			n:        1,
			expected: 0,
		},
		{
			name:     "Example 3",
			n:        4,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minSteps(tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
