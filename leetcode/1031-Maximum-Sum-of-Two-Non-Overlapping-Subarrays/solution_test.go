package leetcode

import "testing"

func TestMaxSumTwoNoOverlap(t *testing.T) {
	testCases := []struct {
		name      string
		nums      []int
		firstLen  int
		secondLen int
		expected  int
	}{
		{
			name:      "Example 1",
			nums:      []int{0, 6, 5, 2, 2, 5, 1, 9, 4},
			firstLen:  1,
			secondLen: 2,
			expected:  20,
		},
		{
			name:      "Example 2",
			nums:      []int{3, 8, 1, 3, 2, 1, 8, 9, 0},
			firstLen:  3,
			secondLen: 2,
			expected:  29,
		},
		{
			name:      "Example 3",
			nums:      []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
			firstLen:  4,
			secondLen: 3,
			expected:  31,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSumTwoNoOverlap(tc.nums, tc.firstLen, tc.secondLen)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
