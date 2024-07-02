package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	if min := minStack.GetMin(); min != -3 {
		t.Errorf("GetMin() = %d; want -3", min)
	}

	minStack.Pop()

	if top := minStack.Top(); top != 0 {
		t.Errorf("Top() = %d; want 0", top)
	}

	if min := minStack.GetMin(); min != -2 {
		t.Errorf("GetMin() = %d; want -2", min)
	}
}
