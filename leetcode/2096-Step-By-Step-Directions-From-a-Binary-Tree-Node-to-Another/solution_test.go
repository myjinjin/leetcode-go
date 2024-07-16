package leetcode

import "testing"

func TestGetDirections(t *testing.T) {
	testCases := []struct {
		name       string
		root       *TreeNode
		startValue int
		destValue  int
		expected   string
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 4}},
			},
			startValue: 3,
			destValue:  6,
			expected:   "UURL",
		},
		{
			name: "Example 2",
			root: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1},
			},
			startValue: 2,
			destValue:  1,
			expected:   "L",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getDirections(tc.root, tc.startValue, tc.destValue)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
