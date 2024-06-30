package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var cycleNode *ListNode
	curr := head
	nodeMap := make(map[*ListNode]bool)
	for curr != nil {
		if _, ok := nodeMap[curr]; ok {
			cycleNode = curr
			break
		} else {
			nodeMap[curr] = true
		}
		curr = curr.Next
	}

	return cycleNode
}
