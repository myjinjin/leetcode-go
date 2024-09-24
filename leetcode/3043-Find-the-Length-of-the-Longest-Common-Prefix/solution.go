package leetcode

import "strconv"

type trieNode struct {
	children [10]*trieNode
	isEnd    bool
}

func (n *trieNode) insert(str string) {
	curr := n
	for i, c := range str {
		index := int(c - '0')
		if curr.children[index] == nil {
			curr.children[index] = &trieNode{}
		}
		curr = curr.children[index]
		if i == len(str)-1 {
			curr.isEnd = true
		}
	}
}

func (n *trieNode) findPrefixLength(str string) int {
	curr := n
	for i, c := range str {
		index := int(c - '0')
		if curr.children[index] == nil {
			return i
		}
		curr = curr.children[index]
	}
	return len(str)
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	longestLength := 0
	root := &trieNode{}

	for _, n := range arr1 {
		root.insert(strconv.Itoa(n))
	}

	for _, n := range arr2 {
		length := root.findPrefixLength(strconv.Itoa(n))
		if length > longestLength {
			longestLength = length
		}
	}

	return longestLength
}
