package leetcode

import "testing"

func TestMaxProfitAssignment(t *testing.T) {
	testCases := []struct {
		name       string
		difficulty []int
		profit     []int
		worker     []int
		expected   int
	}{
		{
			name:       "Example 1",
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			expected:   100,
		},
		{
			name:       "Example 2",
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			expected:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProfitAssignment(tc.difficulty, tc.profit, tc.worker)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
