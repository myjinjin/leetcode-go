package leetcode

import "testing"

func TestTotalCost(t *testing.T) {
	testCases := []struct {
		name       string
		costs      []int
		k          int
		candidates int
		expected   int64
	}{
		{
			name:       "Example 1",
			costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
			k:          3,
			candidates: 4,
			expected:   11,
		},
		{
			name:       "Example 2",
			costs:      []int{1, 2, 4, 1},
			k:          3,
			candidates: 3,
			expected:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := totalCost(tc.costs, tc.k, tc.candidates)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}

}
