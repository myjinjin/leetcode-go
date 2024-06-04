package leetcode

import "testing"

func TestLongestZigZag(t *testing.T) {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input:    buildTree([]int{1, -1, 1, 1, 1, -1, -1, 1, 1, -1, 1, -1, -1, -1, 1}),
			expected: 3,
		},
		{
			input:    buildTree([]int{1, 1, 1, -1, 1, -1, -1, 1, 1, -1, 1}),
			expected: 4,
		},
		{
			input:    buildTree([]int{1}),
			expected: 0,
		},
	}

	for _, tc := range testCases {
		output := longestZigZag(tc.input)
		if output != tc.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", tc.input, tc.expected, output)
		}
	}
}

func buildTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != -1 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != -1 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
