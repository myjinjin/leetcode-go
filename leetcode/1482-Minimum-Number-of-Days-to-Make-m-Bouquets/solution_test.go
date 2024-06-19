package leetcode

import "testing"

func TestMinDays(t *testing.T) {
	testCases := []struct {
		name     string
		bloomDay []int
		m        int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        1,
			expected: 3,
		},
		{
			name:     "Example 2",
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        2,
			expected: -1,
		},
		{
			name:     "Example 3",
			bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
			m:        2,
			k:        3,
			expected: 12,
		},
		{
			name:     "Example 4",
			bloomDay: []int{1000000000, 1000000000},
			m:        1,
			k:        1,
			expected: 1000000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minDays(tc.bloomDay, tc.m, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
