package leetcode

import "testing"

func TestIsSubPath(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Example 1",
			head:     &ListNode{},
			root:     &TreeNode{},
			expected: true,
		},
		{
			name:     "Example 2",
			head:     &ListNode{},
			root:     &TreeNode{},
			expected: true,
		},
		{
			name:     "Example 3",
			head:     &ListNode{},
			root:     &TreeNode{},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isSubPath(tc.head, tc.root)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
