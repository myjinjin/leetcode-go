package leetcode

import "testing"

func TestSmallestChair(t *testing.T) {
	testCases := []struct {
		name         string
		times        [][]int
		targetFriend int
		expected     int
	}{
		{
			name:         "Example 1",
			times:        [][]int{{1, 4}, {2, 3}, {4, 6}},
			targetFriend: 1,
			expected:     1,
		},
		{
			name:         "Example 2",
			times:        [][]int{{3, 10}, {1, 5}, {2, 6}},
			targetFriend: 0,
			expected:     2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := smallestChair(tc.times, tc.targetFriend)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
