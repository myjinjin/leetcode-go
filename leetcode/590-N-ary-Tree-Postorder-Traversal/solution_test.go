package leetcode

import (
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	testCases := []struct {
		name     string
		root     *Node
		expected []int
	}{
		{
			name: "Example 1",
			root: &Node{
				Val: 1,
				Children: []*Node{
					{Val: 3, Children: []*Node{
						{Val: 5},
						{Val: 6},
					}},
					{Val: 2},
					{Val: 4},
				},
			},
			expected: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name: "Example 2",
			root: &Node{
				Val: 1,
				Children: []*Node{
					{Val: 2},
					{Val: 3, Children: []*Node{
						{Val: 6},
						{Val: 7, Children: []*Node{{Val: 11, Children: []*Node{{Val: 14}}}}},
					}},
					{Val: 4, Children: []*Node{
						{Val: 8, Children: []*Node{{Val: 12}}},
					}},
					{Val: 5, Children: []*Node{
						{Val: 9, Children: []*Node{{Val: 13}}},
						{Val: 10},
					}},
				},
			},
			expected: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := postorder(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
