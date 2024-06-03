# [2130. Maximum Twin Sum of a Linked List](https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/)

## Problem

In a linked list of size `n`, where `n` is even, the `ith` node (0-indexed) of the linked list is known as the twin of the `(n-1-i)th` node, if `0 <= i <= (n / 2) - 1`.

For example, if `n = 4`, then node `0` is the twin of node `3`, and node `1` is the twin of node `2`. These are the only nodes with twins for `n = 4`.
The twin sum is defined as the sum of a node and its twin.

Given the `head` of a linked list with even length, return the maximum twin sum of the linked list.

Example 1:

![alt text](image.png)

```
Input: head = [5,4,2,1]
Output: 6
Explanation:
Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
There are no other nodes with twins in the linked list.
Thus, the maximum twin sum of the linked list is 6. 
```

Example 2:

![alt text](image-1.png)

```
Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7. 
```

Example 3:

![alt text](image-2.png)

```
Input: head = [1,100000]
Output: 100001
Explanation:
There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.
```

Constraints:

- The number of nodes in the list is an even integer in the range `[2, 10^5]`.
- `1 <= Node.val <= 10^5`

## Solution

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    idx := 0
    value := make(map[int]int)
    for head != nil {
        value[idx] = head.Val
        head = head.Next
        idx++
    }

    n := len(value)
    maxSum := 0

    for i:=0; i<n/2; i++ {
        sum := value[i] + value[n-1-i]
        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}
```


```go
func pairSum(head *ListNode) int {
    // 중간 지점 찾기
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 뒷부분 리스트 뒤집기
    var prev *ListNode
    for slow != nil {
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }

    maxSum := 0
    first, second := head, prev
    for second != nil {
        sum := first.Val + second.Val
        if sum > maxSum {
            maxSum = sum
        }
        first = first.Next
        second = second.Next
    }

    return maxSum
}
```