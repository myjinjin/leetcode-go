package leetcode

import "testing"

func TestCheckInclusion(t *testing.T) {
	testCases := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "Example 2",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkInclusion(tc.s1, tc.s2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
