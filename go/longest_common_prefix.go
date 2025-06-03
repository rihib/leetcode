//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	var currentChar byte
	currentIndex := 0
	for len(strs) > 0 {
		for i, s := range strs {
			if currentIndex >= len(s) {
				return prefix.String()
			}
			if i == 0 {
				currentChar = s[currentIndex]
			}
			if s[currentIndex] != currentChar {
				return prefix.String()
			}
		}
		prefix.WriteByte(currentChar)
		currentIndex++
	}
	return prefix.String()
}

/*
Trie
*/
type trieNode struct {
	children  map[rune]*trieNode
	isWordEnd bool
}

func newTrie() *trieNode {
	return &trieNode{make(map[rune]*trieNode), false}
}

func (t *trieNode) insert(word string) {
	node := t
	for _, r := range word {
		next, ok := node.children[r]
		if !ok {
			node.children[r] = &trieNode{make(map[rune]*trieNode), false}
			next = node.children[r]
		}
		node = next
	}
	node.isWordEnd = true
}

func (t *trieNode) commonPrefix() string {
	var prefix strings.Builder
	node := t
	for node != nil && len(node.children) == 1 && !node.isWordEnd {
		for r, next := range node.children {
			prefix.WriteRune(r)
			node = next
		}
	}
	return prefix.String()
}

func longestCommonPrefixTrie(strs []string) string {
	t := newTrie()
	for _, s := range strs {
		t.insert(s)
	}
	return t.commonPrefix()
}
