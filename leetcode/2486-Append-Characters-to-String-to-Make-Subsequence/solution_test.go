package leetcode

import "testing"

func TestAppendCharacters(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "coaching",
			t:        "coding",
			expected: 4,
		},
		{
			s:        "abcde",
			t:        "a",
			expected: 0,
		},
		{
			s:        "z",
			t:        "abcde",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		result := appendCharacters(tc.s, tc.t)
		if result != tc.expected {
			t.Errorf("appendCharacters(%q, %q) = %d; expected %d", tc.s, tc.t, result, tc.expected)
		}
	}
}
