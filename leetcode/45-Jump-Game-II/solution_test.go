package leetcode

import "testing"

func TestJump(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := jump(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
