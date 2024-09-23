package leetcode

import "testing"

func TestMinExtraChar(t *testing.T) {
	testCases := []struct {
		name       string
		s          string
		dictionary []string
		expected   int
	}{
		{
			name:       "Example 1",
			s:          "leetscode",
			dictionary: []string{"leet", "code", "leetcode"},
			expected:   1,
		},
		{
			name:       "Example 2",
			s:          "sayhelloworld",
			dictionary: []string{"hello", "world"},
			expected:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minExtraChar(tc.s, tc.dictionary)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
