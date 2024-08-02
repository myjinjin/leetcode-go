package leetcode

import "testing"

func TestMaximizeWin(t *testing.T) {
	testCases := []struct {
		name           string
		prizePositions []int
		k              int
		expected       int
	}{
		{
			name:           "Example 1",
			prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
			k:              2,
			expected:       7,
		},
		{
			name:           "Example 2",
			prizePositions: []int{1, 2, 3, 4},
			k:              0,
			expected:       2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximizeWin(tc.prizePositions, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
