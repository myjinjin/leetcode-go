package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	for headA != nil && headB != nil {
		if _, ok := nodeMap[headA]; ok {
			return headA
		}
		nodeMap[headA] = true
		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		nodeMap[headB] = true

		headA = headA.Next
		headB = headB.Next
	}

	for headA != nil {
		if _, ok := nodeMap[headA]; ok {
			return headA
		}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
