package leetcode

import "testing"

func TestMinIncrementForUnique(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 2},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 2, 1, 7},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		result := minIncrementForUnique(tc.nums)
		if result != tc.expected {
			t.Errorf("Expected: %v, but got %v", tc.expected, result)
		}
	}
}
