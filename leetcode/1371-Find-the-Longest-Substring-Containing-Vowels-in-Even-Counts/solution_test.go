package leetcode

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "eleetminicoworoep",
			expected: 13,
		},
		{
			name:     "Example 2",
			s:        "leetcodeisgreat",
			expected: 5,
		},
		{
			name:     "Example 3",
			s:        "bcbcbc",
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findTheLongestSubstring(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
