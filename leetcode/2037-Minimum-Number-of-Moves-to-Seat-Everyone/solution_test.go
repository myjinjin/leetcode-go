package leetcode

import "testing"

func TestMinMovesToSeat(t *testing.T) {
	testCases := []struct {
		name     string
		seats    []int
		students []int
		expected int
	}{
		{
			name:     "Example 1",
			seats:    []int{3, 1, 5},
			students: []int{2, 7, 4},
			expected: 4,
		},
		{
			name:     "Example 2",
			seats:    []int{4, 1, 5, 9},
			students: []int{1, 3, 2, 6},
			expected: 7,
		},
		{
			name:     "Example 3",
			seats:    []int{2, 2, 6, 6},
			students: []int{1, 3, 2, 6},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minMovesToSeat(tc.seats, tc.students)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
