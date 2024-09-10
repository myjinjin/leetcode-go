package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		if curr.Next != nil {
			n := gcd(curr.Next.Val, curr.Val)
			newNode := &ListNode{Val: n, Next: curr.Next}
			curr.Next = newNode
			curr = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
