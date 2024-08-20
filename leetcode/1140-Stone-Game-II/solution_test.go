package leetcode

import "testing"

func TestStoneGameII(t *testing.T) {
	testCases := []struct {
		name     string
		piles    []int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{2, 7, 9, 4, 4},
			expected: 10,
		},
		{
			name:     "Example 2",
			piles:    []int{1, 2, 3, 4, 5, 100},
			expected: 104,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := stoneGameII(tc.piles)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
