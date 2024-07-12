package leetcode

import "testing"

func TestMaximumGain(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		x        int
		y        int
		expected int
	}{
		{
			name:     "Example 1",
			s:        "cdbcbbaaabab",
			x:        4,
			y:        5,
			expected: 19,
		},
		{
			name:     "Example 2",
			s:        "aabbaaxybbaabb",
			x:        5,
			y:        4,
			expected: 20,
		},
		{
			name:     "Example 3",
			s:        "aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha",
			x:        1926,
			y:        4320,
			expected: 112374,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumGain(tc.s, tc.x, tc.y)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
