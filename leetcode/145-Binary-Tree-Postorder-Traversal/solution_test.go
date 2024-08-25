package leetcode

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Example 1",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Example 3",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := postorderTraversal(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
