package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		name     string
		preorder []int
		inorder  []int
		expected *TreeNode
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
		},
		{
			name:     "Example 2",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: &TreeNode{Val: -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := buildTree(tc.preorder, tc.inorder)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
