package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	testCases := []struct {
		name      string
		cpdomains []string
		expected  []string
	}{
		{
			name:      "Example 1",
			cpdomains: []string{"9001 discuss.leetcode.com"},
			expected:  []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"},
		},
		{
			name:      "Example 2",
			cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			expected:  []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subdomainVisits(tc.cpdomains)
			sort.Strings(result)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
