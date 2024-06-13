package leetcode

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "Example 2",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minCostClimbingStairs(tc.cost)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

// func minCostClimbingStairs(cost []int) int {

// }
