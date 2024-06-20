package leetcode

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	testCases := []struct {
		name       string
		products   []string
		searchWord string
		expected   [][]string
	}{
		{
			name:       "Example 1",
			products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			searchWord: "mouse",
			expected:   [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
		{
			name:       "Example 2",
			products:   []string{"havana"},
			searchWord: "havana",
			expected:   [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := suggestedProducts(tc.products, tc.searchWord)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
