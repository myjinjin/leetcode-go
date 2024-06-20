package leetcode

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := [][]string{}

	root := &Trie{children: make(map[rune]*Trie)}

	for _, product := range products {
		curr := root
		for _, ch := range product {
			if curr.children[ch] == nil {
				curr.children[ch] = &Trie{children: make(map[rune]*Trie)}
			}
			curr = curr.children[ch]
			curr.words = append(curr.words, product)
		}
	}

	node := root
	for _, ch := range searchWord {
		if node != nil {
			node = node.children[ch]
		}
		if node == nil {
			result = append(result, []string{})
		} else {
			if len(node.words) <= 3 {
				result = append(result, node.words)
			} else {
				result = append(result, node.words[:3])
			}
		}
	}

	return result
}

type Trie struct {
	children map[rune]*Trie
	words    []string
}
