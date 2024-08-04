package leetcode

import "testing"

func TestClosestCost(t *testing.T) {
	testCases := []struct {
		name         string
		baseCosts    []int
		toppingCosts []int
		target       int
		expected     int
	}{
		{
			name:         "Example 1",
			baseCosts:    []int{1, 7},
			toppingCosts: []int{3, 4},
			target:       10,
			expected:     10,
		},
		{
			name:         "Example 2",
			baseCosts:    []int{2, 3},
			toppingCosts: []int{4, 5, 100},
			target:       18,
			expected:     17,
		},
		{
			name:         "Example 3",
			baseCosts:    []int{3, 10},
			toppingCosts: []int{2, 5},
			target:       9,
			expected:     8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := closestCost(tc.baseCosts, tc.toppingCosts, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
