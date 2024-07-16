package leetcode

import "testing"

func TestReverseParentheses(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "(abcd)",
			expected: "dcba",
		},
		{
			name:     "Example 2",
			s:        "(u(love)i)",
			expected: "iloveu",
		},
		{
			name:     "Example 3",
			s:        "(ed(et(oc))el)",
			expected: "leetcode",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseParentheses(tc.s)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
