package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	for _, value := range anagramMap {
		anagrams = append(anagrams, value)
	}

	return anagrams
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	var b strings.Builder
	b.Grow(len(s))
	for _, ch := range r {
		b.WriteRune(ch)
	}
	return b.String()
}
