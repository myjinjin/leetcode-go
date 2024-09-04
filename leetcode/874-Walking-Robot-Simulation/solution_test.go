package leetcode

import "testing"

func TestRobotSim(t *testing.T) {
	testCases := []struct {
		name      string
		commands  []int
		obstacles [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			expected:  25,
		},
		{
			name:      "Example 2",
			commands:  []int{4, -1, 4, -2, 4},
			obstacles: [][]int{{2, 4}},
			expected:  65,
		},
		{
			name:      "Example 3",
			commands:  []int{6, -1, -1, 6},
			obstacles: [][]int{},
			expected:  36,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := robotSim(tc.commands, tc.obstacles)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
