package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
			},
			expected: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}},
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: nil,
		},
		{
			name:     "Example 3",
			root:     &TreeNode{Val: 0},
			expected: &TreeNode{Val: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flatten(tc.root)
			if !reflect.DeepEqual(tc.root, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, tc.root)
			}
		})
	}
}
