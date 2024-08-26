package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	path := []int{}
	helper(root, &path)
	return path
}

func helper(node *Node, path *[]int) {
	if node == nil {
		return
	}

	for _, c := range node.Children {
		helper(c, path)
	}
	*path = append(*path, node.Val)
}
