package leetcode

import "testing"

func TestFindMinDifference(t *testing.T) {
	testCases := []struct {
		name       string
		timePoints []string
		expected   int
	}{
		{
			name:       "Example 1",
			timePoints: []string{"23:59", "00:00"},
			expected:   1,
		},
		{
			name:       "Example 2",
			timePoints: []string{"00:00", "23:59", "00:00"},
			expected:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMinDifference(tc.timePoints)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
