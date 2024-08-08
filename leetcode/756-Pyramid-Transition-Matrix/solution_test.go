package leetcode

import "testing"

func TestPyramidTransition(t *testing.T) {
	testCases := []struct {
		name     string
		bottom   string
		allowed  []string
		expected bool
	}{
		{
			name:     "Example 1",
			bottom:   "BCD",
			allowed:  []string{"BCC", "CDE", "CEA", "FFF"},
			expected: true,
		},
		{
			name:     "Example 2",
			bottom:   "AAAA",
			allowed:  []string{"AAB", "AAC", "BCD", "BBE", "DEF"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := pyramidTransition(tc.bottom, tc.allowed)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
