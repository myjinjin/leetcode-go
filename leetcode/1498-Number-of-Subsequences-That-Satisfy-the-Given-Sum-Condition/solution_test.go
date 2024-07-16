package leetcode

import "testing"

func TestNumSubseq(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 5, 6, 7},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 3, 6, 8},
			target:   10,
			expected: 6,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 3, 3, 4, 6, 7},
			target:   12,
			expected: 61,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numSubseq(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
