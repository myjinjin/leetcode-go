package leetcode

import "testing"

func TestMinSwaps(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "][][",
			expected: 1,
		},
		{
			name:     "Example 2",
			s:        "]]][[[",
			expected: 2,
		},
		{
			name:     "Example 3",
			s:        "[]",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minSwaps(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
