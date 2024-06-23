package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3",
			root:     nil,
			expected: [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := levelOrder(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
