package leetcode

type trieNode struct {
	children map[rune]*trieNode
	count    int
}

func NewTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
		count:    0,
	}
}

func sumPrefixScores(words []string) []int {
	n := len(words)
	result := make([]int, n)

	root := NewTrieNode()
	for _, word := range words {
		node := root
		for _, ch := range word {
			if _, exist := node.children[ch]; !exist {
				node.children[ch] = NewTrieNode()
			}
			node = node.children[ch]
			node.count++
		}
	}

	for i, word := range words {
		node := root
		score := 0
		for _, ch := range word {
			node = node.children[ch]
			score += node.count
		}
		result[i] = score
	}

	return result
}
