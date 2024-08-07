package leetcode

import "testing"

func TestNumberToWords(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected string
	}{
		{
			name:     "Example 1",
			num:      123,
			expected: "One Hundred Twenty Three",
		},
		{
			name:     "Example 2",
			num:      12345,
			expected: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name:     "Example 3",
			num:      1234567,
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name:     "Example 4",
			num:      1000,
			expected: "One Thousand",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numberToWords(tc.num)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
