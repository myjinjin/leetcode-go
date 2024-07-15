package leetcode

import (
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	testCases := []struct {
		name         string
		descriptions [][]int
		expected     *TreeNode
	}{
		{
			name:         "Example 1",
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			expected: &TreeNode{Val: 50,
				Left:  &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 17}},
				Right: &TreeNode{Val: 80, Left: &TreeNode{Val: 19}},
			},
		},
		{
			name:         "Example 2",
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			expected:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := createBinaryTree(tc.descriptions)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
