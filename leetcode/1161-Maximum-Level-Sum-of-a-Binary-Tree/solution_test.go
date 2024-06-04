package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 989,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val: 98693,
					},
					Right: &TreeNode{
						Val: -89388,
						Right: &TreeNode{
							Val: -32127,
						},
					},
				},
			},
			expected: 2,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			result := maxLevelSum(tc.root)
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}
