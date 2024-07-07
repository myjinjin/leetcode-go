package leetcode

import "testing"

func TestNumWaterBottles(t *testing.T) {
	testCases := []struct {
		name        string
		numBottles  int
		numExchange int
		expected    int
	}{
		{
			name:        "Example 1",
			numBottles:  9,
			numExchange: 3,
			expected:    13,
		},
		{
			name:        "Example 2",
			numBottles:  15,
			numExchange: 4,
			expected:    19,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numWaterBottles(tc.numBottles, tc.numExchange)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
