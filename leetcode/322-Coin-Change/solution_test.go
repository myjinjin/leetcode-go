package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := coinChange(tc.coins, tc.amount)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
