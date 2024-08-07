package leetcode

import "testing"

func TestCountPairs(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		distance int
		expected int
	}{
		{
			name:     "Example 1",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
			distance: 3,
			expected: 1,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			distance: 3,
			expected: 2,
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}},
			},
			distance: 3,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countPairs(tc.root, tc.distance)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
