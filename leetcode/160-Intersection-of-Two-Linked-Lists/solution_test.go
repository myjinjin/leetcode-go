package leetcode

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	testCases := []struct {
		name     string
		headA    *ListNode
		headB    *ListNode
		expected *ListNode
	}{
		func() struct {
			name                   string
			headA, headB, expected *ListNode
		} {
			headA, headB, expected := createIntersectedLists(
				[]int{4, 1, 8, 4, 5},
				[]int{5, 6, 1, 8, 4, 5},
				2, 3)
			return struct {
				name                   string
				headA, headB, expected *ListNode
			}{
				name:     "Example 1",
				headA:    headA,
				headB:    headB,
				expected: expected,
			}
		}(),
		func() struct {
			name                   string
			headA, headB, expected *ListNode
		} {
			headA, headB, expected := createIntersectedLists(
				[]int{1, 9, 1, 2, 4},
				[]int{3, 2, 4},
				3, 1)
			return struct {
				name                   string
				headA, headB, expected *ListNode
			}{
				name:     "Example 2",
				headA:    headA,
				headB:    headB,
				expected: expected,
			}
		}(),
		{
			name:     "Example 3",
			headA:    &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			headB:    &ListNode{Val: 1, Next: &ListNode{Val: 5}},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getIntersectionNode(tc.headA, tc.headB)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
func createIntersectedLists(listA, listB []int, skipA, skipB int) (*ListNode, *ListNode, *ListNode) {
	buildList := func(values []int) *ListNode {
		dummy := &ListNode{}
		current := dummy
		for _, val := range values {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		return dummy.Next
	}

	headA := buildList(listA)
	headB := buildList(listB)

	intersectA := headA
	for i := 0; i < skipA; i++ {
		intersectA = intersectA.Next
	}
	intersectB := headB
	for i := 0; i < skipB-1; i++ {
		intersectB = intersectB.Next
	}
	intersectB.Next = intersectA

	return headA, headB, intersectA
}
