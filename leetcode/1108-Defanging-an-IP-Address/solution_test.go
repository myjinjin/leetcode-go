package leetcode

import "testing"

func TestDefangIPaddr(t *testing.T) {
	testCases := []struct {
		name     string
		address  string
		expected string
	}{
		{
			name:     "Example 1",
			address:  "1.1.1.1",
			expected: "1[.]1[.]1[.]1",
		},
		{
			name:     "Example 2",
			address:  "255.100.50.0",
			expected: "255[.]100[.]50[.]0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := defangIPaddr(tc.address)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
