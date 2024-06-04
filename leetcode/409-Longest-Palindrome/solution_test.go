package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abccccdd",
			expected: 7,
		},
		{
			input:    "a",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}
