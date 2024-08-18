package leetcode

import "testing"

func TestMinProcessingTime(t *testing.T) {
	testCases := []struct {
		name          string
		processorTime []int
		tasks         []int
		expected      int
	}{
		{
			name:          "Example 1",
			processorTime: []int{8, 10},
			tasks:         []int{2, 2, 3, 1, 8, 7, 4, 5},
			expected:      16,
		},
		{
			name:          "Example 2",
			processorTime: []int{10, 20},
			tasks:         []int{2, 3, 1, 2, 5, 8, 4, 3},
			expected:      23,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minProcessingTime(tc.processorTime, tc.tasks)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
