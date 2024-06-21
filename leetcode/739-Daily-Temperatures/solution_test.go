package leetcode

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "Example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example 2",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "Example 3",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := dailyTemperatures(tc.temperatures)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
