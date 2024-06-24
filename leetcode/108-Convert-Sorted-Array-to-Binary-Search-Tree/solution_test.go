package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected *TreeNode
	}{
		{
			name: "Example 1",
			nums: []int{-10, -3, 0, 5, 9},
			expected: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}},
				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}},
			},
		},
		{
			name: "Example 2",
			nums: []int{1, 3},
			expected: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sortedArrayToBST(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
