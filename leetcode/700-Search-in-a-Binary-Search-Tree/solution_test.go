package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		val      int
		expected *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val: 2,
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val:      5,
			expected: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			result := searchBST(tc.root, tc.val)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
