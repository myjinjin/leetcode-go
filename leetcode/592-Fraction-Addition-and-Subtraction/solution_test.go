package leetcode

import "testing"

func TestFractionAddition(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		expected   string
	}{
		{
			name:       "Example 1",
			expression: "-1/2+1/2",
			expected:   "0/1",
		},
		{
			name:       "Example 2",
			expression: "-1/2+1/2+1/3",
			expected:   "1/3",
		},
		{
			name:       "Example 3",
			expression: "1/3-1/2",
			expected:   "-1/6",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fractionAddition(tc.expression)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
