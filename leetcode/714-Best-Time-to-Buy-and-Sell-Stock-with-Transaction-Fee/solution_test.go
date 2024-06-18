package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		fee      int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{1, 3, 2, 8, 4, 9},
			fee:      2,
			expected: 8,
		},
		{
			name:     "Example 2",
			prices:   []int{1, 3, 7, 5, 10, 3},
			fee:      3,
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProfit(tc.prices, tc.fee)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
