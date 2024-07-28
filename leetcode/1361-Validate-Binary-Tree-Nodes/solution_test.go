package leetcode

import (
	"testing"
)

func TestValidateBinaryTreeNodes(t *testing.T) {
	testCases := []struct {
		name       string
		n          int
		leftChild  []int
		rightChild []int
		expected   bool
	}{
		{
			name:       "Example 1",
			n:          4,
			leftChild:  []int{1, -1, 3, -1},
			rightChild: []int{2, -1, -1, -1},
			expected:   true,
		},
		{
			name:       "Example 2",
			n:          4,
			leftChild:  []int{1, -1, 3, -1},
			rightChild: []int{2, 3, -1, -1},
			expected:   false,
		},
		{
			name:       "Example 3",
			n:          2,
			leftChild:  []int{1, 0},
			rightChild: []int{-1, -1},
			expected:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validateBinaryTreeNodes(tc.n, tc.leftChild, tc.rightChild)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
