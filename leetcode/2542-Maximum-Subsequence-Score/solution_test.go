package leetcode

import "testing"

func TestMaxScore(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 3, 3, 2},
			nums2:    []int{2, 1, 3, 4},
			k:        3,
			expected: 12,
		},
		{
			name:     "Example 2",
			nums1:    []int{4, 2, 3, 1, 1},
			nums2:    []int{7, 5, 10, 9, 6},
			k:        1,
			expected: 30,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxScore(tc.nums1, tc.nums2, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
