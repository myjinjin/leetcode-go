# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## Problem

Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.


Example 1:

![alt text](image.png)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

Example 2:

```
Input: head = [1], n = 1
Output: []
```

Example 3:

```
Input: head = [1,2], n = 1
Output: [1]
``` 

Constraints:

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## Solution

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := 0
	node := head
	for node != nil {
		sz++
		node = node.Next
	}

	if sz == n {
		head = head.Next
		return head
	}

	nth := head
	prev := &ListNode{}
	for i := 0; i < sz-n; i++ {
		if i == sz-n-1 {
			prev = nth
		}
		nth = nth.Next
	}

	prev.Next = nth.Next
	return head
}
```
