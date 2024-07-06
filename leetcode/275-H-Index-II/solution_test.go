package leetcode

import "testing"

func TestHIndex(t *testing.T) {
	testCases := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "Example 1",
			citations: []int{0, 1, 3, 5, 6},
			expected:  3,
		},
		{
			name:      "Example 2",
			citations: []int{1, 2, 100},
			expected:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hIndex(tc.citations)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
