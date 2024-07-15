package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	childSet := make(map[int]bool)

	for _, d := range descriptions {
		parent, child, left := d[0], d[1], d[2]

		if _, exist := nodeMap[parent]; !exist {
			nodeMap[parent] = &TreeNode{Val: parent}
		}
		if _, childExist := nodeMap[child]; !childExist {
			nodeMap[child] = &TreeNode{Val: child}
		}

		parentNode := nodeMap[parent]
		childNode := nodeMap[child]

		if left == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		childSet[child] = true
	}

	var head *TreeNode
	for val, node := range nodeMap {
		if !childSet[val] {
			head = node
			break
		}
	}
	return head
}
