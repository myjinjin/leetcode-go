package leetcode

import "testing"

func TestMaxKelements(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			nums:     []int{10, 10, 10, 10, 10},
			k:        5,
			expected: 50,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 10, 3, 3, 3},
			k:        3,
			expected: 17,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxKelements(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
