package leetcode

import "testing"

func TestFlipEquiv(t *testing.T) {
	testCases := []struct {
		name     string
		root1    []interface{}
		root2    []interface{}
		expected bool
	}{
		{
			name:     "Example 1",
			root1:    []interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8},
			root2:    []interface{}{1, 3, 2, nil, 6, 4, 5, nil, nil, nil, nil, 8, 7},
			expected: true,
		},
		{
			name:     "Example 2",
			root1:    []interface{}{},
			root2:    []interface{}{},
			expected: true,
		},
		{
			name:     "Example 3",
			root1:    []interface{}{},
			root2:    []interface{}{1},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree1 := createTree(tc.root1)
			tree2 := createTree(tc.root2)

			result := flipEquiv(tree1, tree2)

			if result != tc.expected {
				t.Errorf("Test case %s failed: expected %v, got %v",
					tc.name, tc.expected, result)
			}
		})
	}
}

func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
