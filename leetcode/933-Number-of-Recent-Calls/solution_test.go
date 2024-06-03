package leetcode

import "testing"

func TestRecentCounter(t *testing.T) {
	recentCounter := Constructor()

	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{100, 2},
		{3001, 3},
		{3002, 3},
	}

	for _, tc := range testCases {
		result := recentCounter.Ping(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %d, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}
