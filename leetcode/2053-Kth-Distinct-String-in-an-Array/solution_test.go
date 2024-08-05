package leetcode

import "testing"

func TestKthDistinct(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []string
		k        int
		expected string
	}{
		{
			name:     "Example 1",
			arr:      []string{"d", "b", "c", "b", "c", "a"},
			k:        2,
			expected: "a",
		},
		{
			name:     "Example 2",
			arr:      []string{"aaa", "aa", "a"},
			k:        1,
			expected: "aaa",
		},
		{
			name:     "Example 3",
			arr:      []string{"a", "b", "a"},
			k:        3,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := kthDistinct(tc.arr, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
