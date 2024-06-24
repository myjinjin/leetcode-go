package leetcode

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3},
			},
			expected: 3,
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := diameterOfBinaryTree(tc.root)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
