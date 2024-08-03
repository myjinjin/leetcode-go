package leetcode

import "testing"

func TestCanBeEqual(t *testing.T) {
	testCases := []struct {
		name     string
		target   []int
		arr      []int
		expected bool
	}{
		{
			name:     "Example 1",
			target:   []int{1, 2, 3, 4},
			arr:      []int{2, 4, 1, 3},
			expected: true,
		},
		{
			name:     "Example 2",
			target:   []int{7},
			arr:      []int{7},
			expected: true,
		},
		{
			name:     "Example 3",
			target:   []int{3, 7, 9},
			arr:      []int{3, 7, 11},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canBeEqual(tc.target, tc.arr)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
