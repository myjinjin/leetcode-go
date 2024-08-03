package leetcode

import "testing"

func TestCanReach(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		minJump  int
		maxJump  int
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "011010",
			minJump:  2,
			maxJump:  3,
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "01101110",
			minJump:  2,
			maxJump:  3,
			expected: false,
		},
		{
			name:     "Example 3",
			s:        "0000000000",
			minJump:  8,
			maxJump:  8,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canReach(tc.s, tc.minJump, tc.maxJump)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
