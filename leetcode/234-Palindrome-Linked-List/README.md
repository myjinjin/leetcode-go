# [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

## Problem

Given the `head` of a singly linked list, return `true` if it is a **palindrome** or `false` otherwise.


Example 1:

![alt text](image-1.png)

```
Input: head = [1,2,2,1]
Output: true
```

Example 2:

![alt text](image.png)

```
Input: head = [1,2]
Output: false
```

Constraints:

- The number of nodes in the list is in the range `[1, 10^5]`.
- `0 <= Node.val <= 9`

## Solution

```go
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	rightList := reverseList(slow.Next)

	leftList := head
	for rightList != nil {
		if leftList.Val != rightList.Val {
			return false
		}
		leftList = leftList.Next
		rightList = rightList.Next
	}

	return true
}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
```