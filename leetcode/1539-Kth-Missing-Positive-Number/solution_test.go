package leetcode

import "testing"

func TestFindKthPositive(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		k        int
		expected int
	}{
		{
			name:     "Exmample 1",
			arr:      []int{2, 3, 4, 7, 11},
			k:        5,
			expected: 9,
		},
		{
			name:     "Exmample 2",
			arr:      []int{1, 2, 3, 4},
			k:        2,
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findKthPositive(tc.arr, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
