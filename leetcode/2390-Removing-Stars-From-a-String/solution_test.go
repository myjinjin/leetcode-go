package leetcode

import "testing"

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "leet**cod*e",
			expected: "lecoe",
		},
		{
			input:    "erase*****",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := removeStars(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
