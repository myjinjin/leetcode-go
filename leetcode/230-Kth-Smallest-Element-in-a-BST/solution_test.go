package leetcode

import "testing"

func TestKthSmallest(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		k        int
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4},
			},
			k:        1,
			expected: 1,
		},
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:        3,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := kthSmallest(tc.root, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
