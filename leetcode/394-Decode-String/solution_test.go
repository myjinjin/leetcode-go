package leetcode

import "testing"

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
	}

	for _, tc := range testCases {
		result := decodeString(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
		}
	}
}
