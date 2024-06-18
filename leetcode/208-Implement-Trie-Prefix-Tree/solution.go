package leetcode

type Node struct {
	children    map[rune]*Node
	isEndOfWord bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children:    make(map[rune]*Node),
			isEndOfWord: false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &Node{
				children:    make(map[rune]*Node),
				isEndOfWord: false,
			}
		}
		node = node.children[ch]
	}
	node.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}
