package leetcode

import (
	"reflect"
	"testing"
)

func TestDelNodes(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		toDelete []int
		expected []*TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			toDelete: []int{3, 5},
			expected: []*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}},
				{Val: 6},
				{Val: 7},
			},
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
			toDelete: []int{3},
			expected: []*TreeNode{{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := delNodes(tc.root, tc.toDelete)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
