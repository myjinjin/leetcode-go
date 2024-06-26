package leetcode

import (
	"reflect"
	"testing"
)

func TestBalanceBST(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			},
		},
		{
			name:     "Example 2",
			root:     &TreeNode{},
			expected: &TreeNode{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := balanceBST(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
