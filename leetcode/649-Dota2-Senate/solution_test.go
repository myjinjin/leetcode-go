package leetcode

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "RD",
			expected: "Radiant",
		},
		{
			input:    "RDD",
			expected: "Dire",
		},
	}

	for _, tc := range testCases {
		result := predictPartyVictory(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
		}
	}
}
