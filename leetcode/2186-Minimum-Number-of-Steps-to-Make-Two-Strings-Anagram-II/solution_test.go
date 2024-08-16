package leetcode

import "testing"

func TestMinSteps(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			t:        "coats",
			expected: 7,
		},
		{
			name:     "Example 2",
			s:        "night",
			t:        "thing",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minSteps(tc.s, tc.t)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
