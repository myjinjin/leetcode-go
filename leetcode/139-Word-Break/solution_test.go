package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := wordBreak(tc.s, tc.wordDict)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
