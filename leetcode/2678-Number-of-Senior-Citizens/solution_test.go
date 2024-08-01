package leetcode

import "testing"

func TestCountSeniors(t *testing.T) {
	testCases := []struct {
		name     string
		details  []string
		expected int
	}{
		{
			name:     "Example 1",
			details:  []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			expected: 2,
		},
		{
			name:     "Example 2",
			details:  []string{"1313579440F2036", "2921522980M5644"},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countSeniors(tc.details)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
