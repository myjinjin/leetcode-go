package leetcode

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
			},
			expected: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
			},
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
		},
		{
			name:     "Example 3",
			root:     nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := invertTree(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
