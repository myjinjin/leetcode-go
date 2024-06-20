package leetcode

import "testing"

func TestMaxDistance(t *testing.T) {
	testCases := []struct {
		name     string
		position []int
		m        int
		expected int
	}{
		{
			name:     "Example 1",
			position: []int{1, 2, 3, 4, 7},
			m:        3,
			expected: 3,
		},
		{
			name:     "Example 2",
			position: []int{5, 4, 3, 2, 1, 1000000000},
			m:        2,
			expected: 999999999,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxDistance(tc.position, tc.m)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
