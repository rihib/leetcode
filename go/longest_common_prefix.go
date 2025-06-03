//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix strings.Builder
	currentIndex := 0
	for {
		var currentChar byte
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
}

/*
Trie
*/
type Trie struct {
	root *Node
}

type Node struct {
	children  map[rune]*Node
	isWordEnd bool
}

func NewTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[r]
	}
	node.isWordEnd = true
}

func (t *Trie) Prefix() string {
	var prefix strings.Builder
	node := t.root
	for len(node.children) == 1 && !node.isWordEnd {
		for r, child := range node.children {
			prefix.WriteRune(r)
			node = child // ここでnodeをchildで更新してもrangeは評価済みなのでループが続くことはない
		}
	}
	return prefix.String()
}

func longestCommonPrefixTrie(strs []string) string {
	t := NewTrie()
	for _, word := range strs {
		t.Insert(word)
	}
	return t.Prefix()
}
