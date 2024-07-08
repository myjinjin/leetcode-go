package leetcode

import "testing"

func TestTriangleNumbert(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 2, 3, 4},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := triangleNumber(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
