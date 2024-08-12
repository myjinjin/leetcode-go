package leetcode

import "testing"

func TestMaximumTotalDamage(t *testing.T) {
	testCases := []struct {
		name     string
		power    []int
		expected int64
	}{
		{
			name:     "Example 1",
			power:    []int{1, 1, 3, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			power:    []int{7, 1, 6, 6},
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumTotalDamage(tc.power)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
