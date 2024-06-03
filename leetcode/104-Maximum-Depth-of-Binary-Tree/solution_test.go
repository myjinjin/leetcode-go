package leetcode

import "testing"

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		result := maxDepth(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %v, Expected: %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
