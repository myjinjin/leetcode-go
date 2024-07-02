# [155. Min Stack](https://leetcode.com/problems/min-stack/)

## Problem

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

 

Example 1:

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
```

Constraints:

- `-2^31 <= val <= 2^31 - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on non-empty stacks.
- At most `3 * 10^4` calls will be made to `push`, `pop`, `top`, and `getMin`.

## Solution

```go
type Node struct {
	value int
	min   int
	next  *Node
}

func NewNode(value, min int, next *Node) *Node {
	return &Node{
		value: value,
		min:   min,
		next:  next,
	}
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{head: nil}
}

func (s *MinStack) Push(val int) {
	if s.head == nil {
		s.head = NewNode(val, val, nil)
		return
	}

	min := s.head.min
	if min > val {
		min = val
	}
	s.head = NewNode(val, min, s.head)
}

func (s *MinStack) Pop() {
	s.head = s.head.next
}

func (s *MinStack) Top() int {
	return s.head.value
}

func (s *MinStack) GetMin() int {
	return s.head.min
}
```