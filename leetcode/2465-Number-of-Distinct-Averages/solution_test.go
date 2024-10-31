package leetcode

import "testing"

func TestDistinctAverages(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 1, 4, 0, 3, 5},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 100},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := distinctAverages(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
