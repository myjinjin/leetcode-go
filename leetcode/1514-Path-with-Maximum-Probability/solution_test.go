package leetcode

import "testing"

func TestMaxProbability(t *testing.T) {
	testCases := []struct {
		name      string
		n         int
		edges     [][]int
		succProb  []float64
		startNode int
		endNode   int
		expected  float64
	}{
		{
			name:      "Example 1",
			n:         3,
			edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
			succProb:  []float64{0.5, 0.5, 0.2},
			startNode: 0,
			endNode:   2,
			expected:  0.25,
		},
		{
			name:      "Example 2",
			n:         3,
			edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
			succProb:  []float64{0.5, 0.5, 0.3},
			startNode: 0,
			endNode:   2,
			expected:  0.3,
		},
		{
			name:      "Example 3",
			n:         3,
			edges:     [][]int{{0, 1}},
			succProb:  []float64{0.5},
			startNode: 0,
			endNode:   2,
			expected:  0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProbability(tc.n, tc.edges, tc.succProb, tc.startNode, tc.endNode)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
