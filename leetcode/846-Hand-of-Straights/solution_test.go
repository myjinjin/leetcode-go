package leetcode

import "testing"

func TestIsNStraightHand(t *testing.T) {
	testCases := []struct {
		name      string
		hand      []int
		groupSize int
		expected  bool
	}{
		{
			name:      "Example 1",
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			expected:  true,
		},
		{
			name:      "Example 2",
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isNStraightHand(tc.hand, tc.groupSize)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
