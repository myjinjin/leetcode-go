package leetcode

import "testing"

func TestMinimumSteps(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int64
	}{
		{
			name:     "Example 1",
			s:        "101",
			expected: 1,
		},
		{
			name:     "Example 2",
			s:        "100",
			expected: 2,
		},
		{
			name:     "Example 3",
			s:        "0111",
			expected: 0,
		},
		{
			name:     "Example 4",
			s:        "11000111",
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumSteps(tc.s)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
