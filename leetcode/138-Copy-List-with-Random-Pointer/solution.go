package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copies := make(map[*Node]*Node)
	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if c, ok := copies[node]; ok {
			return c
		}

		c := &Node{Val: node.Val}
		copies[node] = c
		c.Next = clone(node.Next)
		c.Random = clone(node.Random)
		return c
	}
	return clone(head)
}
