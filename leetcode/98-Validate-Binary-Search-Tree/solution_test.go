package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Example 1",
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidBST(tc.root)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
