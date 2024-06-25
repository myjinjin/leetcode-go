package leetcode

import (
	"reflect"
	"testing"
)

func TestBstToGst(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}},
				},
			},
			expected: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   36,
					Left:  &TreeNode{Val: 36},
					Right: &TreeNode{Val: 35, Right: &TreeNode{Val: 33}},
				},
				Right: &TreeNode{
					Val:   21,
					Left:  &TreeNode{Val: 26},
					Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 8}},
				},
			},
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 0, Right: &TreeNode{Val: 1}},
			expected: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := bstToGst(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
