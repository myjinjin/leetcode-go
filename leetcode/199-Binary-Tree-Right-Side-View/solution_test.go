package leetcode

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		expected []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 3},
		},
		{
			root:     nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		result := rightSideView(tc.root)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, result)
		}
	}
}
