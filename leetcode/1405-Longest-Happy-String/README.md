# [1405. Longest Happy String](https://leetcode.com/problems/longest-happy-string/)

## Problem

A string `s` is called happy if it satisfies the following conditions:

- `s` only contains the letters `'a'`, `'b'`, and `'c'`.
- `s` does not contain any of `"aaa"`, `"bbb"`, or `"ccc"` as a substring.
- `s` contains at most `a` occurrences of the letter `'a'`.
- `s` contains at most `b` occurrences of the letter `'b'`.
- `s` contains at most `c` occurrences of the letter `'c'`.

Given three integers `a`, `b`, and `c`, return the **longest possible happy** string. If there are multiple longest happy strings, return any of them. If there is no such string, return the empty string `""`.

A **substring** is a contiguous sequence of characters within a string.


Example 1:

```
Input: a = 1, b = 1, c = 7
Output: "ccaccbcc"
Explanation: "ccbccacc" would also be a correct answer.
```

Example 2:

```
Input: a = 7, b = 1, c = 0
Output: "aabaa"
Explanation: It is the only correct answer in this case.
```

Constraints:

- `0 <= a, b, c <= 100`
- `a + b + c > 0`


## Solution

```go
type CharCount struct {
	char  byte
	count int
}

type PriorityQueue []*CharCount

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*CharCount)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	if a > 0 {
		heap.Push(&pq, &CharCount{'a', a})
	}
	if b > 0 {
		heap.Push(&pq, &CharCount{'b', b})
	}
	if c > 0 {
		heap.Push(&pq, &CharCount{'c', c})
	}

	result := make([]byte, 0, a+b+c)

	for pq.Len() > 0 {
		first := heap.Pop(&pq).(*CharCount)
		if len(result) >= 2 && result[len(result)-1] == first.char && result[len(result)-2] == first.char {
			if pq.Len() == 0 {
				break
			}
			second := heap.Pop(&pq).(*CharCount)
			result = append(result, second.char)
			second.count--
			if second.count > 0 {
				heap.Push(&pq, second)
			}
			heap.Push(&pq, first)
		} else {
			result = append(result, first.char)
			first.count--
			if first.count > 0 {
				heap.Push(&pq, first)
			}
		}
	}

	return string(result)
}
```